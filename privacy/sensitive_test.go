package privacy

import "testing"

var benchmarkWalletGroupSink WalletGroup

type User struct {
	Account string
	Phone   string
	Email   string
}

type Wallet struct {
	Balance int
	User    User
}

type WalletList []Wallet

type UserMap map[string]User

type WalletGroup struct {
	Name    string
	Wallets WalletList
	Index   UserMap
}

func (u User) MakeDesensitize(ctx ViewerContext) any {
	phone := "***"
	if ctx.Level == MaskSuperUser {
		phone = u.Phone
	}

	return User{
		Account: u.Account,
		Phone:   phone,
		Email:   u.Email,
	}
}

func (u User) DesensitizeType() DesensitizeType { return DesTypeObject }

func (w Wallet) MakeDesensitize(ctx ViewerContext) any {
	return Wallet{
		Balance: w.Balance,
		User:    w.User.MakeDesensitize(ctx).(User),
	}
}

func (w Wallet) DesensitizeType() DesensitizeType { return DesTypeObject }

func (ws WalletList) MakeDesensitize(ctx ViewerContext) any {
	out := make(WalletList, len(ws))
	for i := range ws {
		out[i] = ws[i].MakeDesensitize(ctx).(Wallet)
	}
	return out
}

func (ws WalletList) DesensitizeType() DesensitizeType { return DesTypeArray }

func (m UserMap) MakeDesensitize(ctx ViewerContext) any {
	out := make(UserMap, len(m))
	for key, value := range m {
		out[key] = value.MakeDesensitize(ctx).(User)
	}
	return out
}

func (m UserMap) DesensitizeType() DesensitizeType { return DesTypeMap }

func (g WalletGroup) MakeDesensitize(ctx ViewerContext) any {
	return WalletGroup{
		Name:    g.Name,
		Wallets: g.Wallets.MakeDesensitize(ctx).(WalletList),
		Index:   g.Index.MakeDesensitize(ctx).(UserMap),
	}
}

func (g WalletGroup) DesensitizeType() DesensitizeType { return DesTypeObject }

func TestManualDesensitize_RecursiveMask(t *testing.T) {
	u := User{Account: "test", Phone: "123456789", Email: "email@ee.com"}
	group := WalletGroup{
		Name: "group-manual",
		Wallets: WalletList{
			{Balance: 0, User: u},
			{Balance: 8, User: u},
		},
		Index: UserMap{"owner": u},
	}

	masked := group.MakeDesensitize(ViewerContext{Role: []string{"public"}, Level: MaskPublic}).(WalletGroup)
	if masked.Wallets[0].User.Phone != "***" {
		t.Fatalf("expected wallet user phone to be masked, got %s", masked.Wallets[0].User.Phone)
	}
	if masked.Index["owner"].Phone != "***" {
		t.Fatalf("expected map user phone to be masked, got %s", masked.Index["owner"].Phone)
	}
	if group.Wallets.DesensitizeType() != DesTypeArray || group.Index.DesensitizeType() != DesTypeMap {
		t.Fatalf("expected type metadata for array/map")
	}
}

func TestManualDesensitize_LevelAware(t *testing.T) {
	u := User{Account: "test", Phone: "123456789", Email: "email@ee.com"}

	public := u.MakeDesensitize(ViewerContext{Role: []string{"guest"}, Level: MaskPublic}).(User)
	if public.Phone != "***" {
		t.Fatalf("expected public phone to be masked, got %s", public.Phone)
	}

	superUser := u.MakeDesensitize(ViewerContext{Role: []string{"admin"}, Level: MaskSuperUser}).(User)
	if superUser.Phone != u.Phone {
		t.Fatalf("expected super user phone to stay original, got %s", superUser.Phone)
	}
}

func BenchmarkDesensitize_Manual(b *testing.B) {
	u := User{Account: "test", Phone: "123456789", Email: "email@ee.com"}
	group := WalletGroup{
		Name: "group-bench",
		Wallets: WalletList{
			{Balance: 0, User: u},
			{Balance: 8, User: u},
			{Balance: 16, User: u},
			{Balance: 32, User: u},
			{Balance: 64, User: u},
			{Balance: 128, User: u},
		},
		Index: UserMap{
			"owner":   u,
			"auditor": u,
			"ops":     u,
			"finance": u,
		},
	}

	ctx := ViewerContext{Role: []string{"public"}, Level: MaskPublic}
	pre := group.MakeDesensitize(ctx).(WalletGroup)
	if pre.Wallets[0].User.Phone != "***" {
		b.Fatalf("expected benchmark pre-check to be masked")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchmarkWalletGroupSink = group.MakeDesensitize(ctx).(WalletGroup)
	}
	if benchmarkWalletGroupSink.Wallets[0].User.Phone != "***" {
		b.Fatalf("expected benchmark result to stay masked")
	}
}
