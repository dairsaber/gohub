package casbin

import (
	"errors"
	"fmt"
	"gohub/pkg/config"
	"gohub/pkg/helpers"
	"sync"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

// CasbinEforcer casbin的操作对象
type CasbinOperator struct {
	db       *gorm.DB
	enforcer *casbin.Enforcer
	adapter  *gormadapter.Adapter
}

type RequestDef struct {
	Sub string
	Obj string
	Act string
}

type NamedRequestDef struct {
	Type string
	RequestDef
}

var once sync.Once

var casbinInstance *CasbinOperator

// CasbinEnforcer casbin的操作对象
func NewCasbinEnforcer(db *gorm.DB) *CasbinOperator {
	once.Do(func() {
		adapter := createAdapter(db)
		enforcer := createEnforcer(adapter)

		casbinInstance = &CasbinOperator{
			db:       db,
			enforcer: enforcer,
			adapter:  adapter,
		}
	})

	return casbinInstance
}

// 初始化gorm casbin适配器 要在new Enforcer之前初始化
func createAdapter(db *gorm.DB) *gormadapter.Adapter {
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		panic(fmt.Sprintf("failed to create casbin enforcer: %v", err))
	}

	return adapter
}

// 初始化casbin
func createEnforcer(adapter *gormadapter.Adapter) *casbin.Enforcer {
	enforcer, err := casbin.NewEnforcer(config.GetString("casbin.model_path"), adapter)
	if err != nil {
		panic(fmt.Sprintf("failed to create casbin enforcer: %v", err))
	}

	return enforcer
}

// 添加策略
func (co *CasbinOperator) AddPolicy(requestDef *RequestDef) (bool, error) {

	if hasPolicy := co.enforcer.HasPolicy(requestDef.Sub, requestDef.Obj, requestDef.Act); !hasPolicy {
		return co.enforcer.AddPolicy(requestDef.Sub, requestDef.Obj, requestDef.Act)
	}

	return false, errors.New("已经存在这条策略")
}

// 添加策略们
func (co *CasbinOperator) AddPolicies(requestDefs *[]RequestDef) (bool, error) {
	policies := getPoliciesMap(requestDefs)
	return co.enforcer.AddPolicies(policies)
}

// 当前命名策略添加授权规则
func (co *CasbinOperator) AddNamedPolicy(requestDef *NamedRequestDef) (bool, error) {

	if hasPolicy := co.enforcer.HasNamedPolicy(requestDef.Type, requestDef.Sub, requestDef.Obj, requestDef.Act); !hasPolicy {
		return co.enforcer.AddNamedPolicy(requestDef.Type, requestDef.Sub, requestDef.Obj, requestDef.Act)
	}

	return false, errors.New("已经存在这条策略")
}

// 当前命名策略添加授权规则们
func (co *CasbinOperator) AddNamedPolicies(name string, requestDefs *[]RequestDef) (bool, error) {
	policies := getPoliciesMap(requestDefs)
	return co.enforcer.AddNamedPolicies(name, policies)
}

// TODO
func (co *CasbinOperator) AddGroupPolicy() {

}

// TODO
func (co *CasbinOperator) AddNamedGroupPolicy() {

}

// TODO
func (co *CasbinOperator) AddGroupPolicies() {

}

// TODO
func (co *CasbinOperator) AddNamedGroupPolicies() {

}

func getPoliciesMap(requestDefs *[]RequestDef) [][]string {

	result := helpers.Reduce(requestDefs, func(prev [][]string, current RequestDef, _ int) [][]string {
		prev = append(prev, []string{current.Sub, current.Obj, current.Act})
		return prev
	}, [][]string{})

	return result

}
