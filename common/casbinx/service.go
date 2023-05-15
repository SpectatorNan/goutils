package casbinx

import (
	"context"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"sync"
)

type Service struct {
	db        *gorm.DB
	modelPath string
}

func NewCasbinService(db *gorm.DB, modelPath string) Service {
	return Service{
		db:        db,
		modelPath: modelPath,
	}
}

// 更新casbin规则
func (s *Service) UpdateCasbin(ctx context.Context, authorityId string, infos []Info) error {
	s.ClearCasbin(0, authorityId)
	rules := [][]string{}
	for _, v := range infos {
		rules = append(rules, []string{authorityId, v.Path, v.Method})
	}
	if len(rules) < 1 {
		return nil
	}
	e := s.casbin()
	ok, _ := e.AddPolicies(rules)
	if !ok {
		return errors.New("casbin has samp api") //errorc.CasbinSampApiErrCode
	}
	return nil
}

// 更新api casbin 规则
func (s *Service) UpdateCasbinApi(ctx context.Context, oldPath string, newPath string, oldMethod string, newMethod string) error {

	err := s.db.WithContext(ctx).Model(&gormadapter.CasbinRule{}).
		Where("v1 = ? and v2 = ?", oldPath, oldMethod).
		Updates(map[string]interface{}{
			"v1": newPath,
			"v2": newMethod,
		}).Error
	return err
}

// 获取权限列表
func (s *Service) GetPolicyPathByAuthorityId(authorityId string) []Info {
	e := s.casbin()
	list := e.GetFilteredPolicy(0, authorityId)

	pathMaps := make([]Info, 0)
	for _, v := range list {
		pathMaps = append(pathMaps, Info{
			Path:   v[1],
			Method: v[2],
		})
	}
	return pathMaps
}

// 清除权限
func (s *Service) ClearCasbin(v int, p ...string) bool {
	e := s.casbin()
	ok, _ := e.RemoveFilteredPolicy(0, p...)
	return ok
}

/**
casbin rule Persistence
*/
var (
	once     sync.Once
	enforcer *casbin.SyncedEnforcer
)

func (s *Service) casbin() *casbin.SyncedEnforcer {
	once.Do(func() {
		adapter, _ := gormadapter.NewAdapterByDB(s.db)
		enforcer, _ = casbin.NewSyncedEnforcer(s.modelPath, adapter)
	})
	_ = enforcer.LoadPolicy()
	return enforcer
}
