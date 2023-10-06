package util

func GOTM_getAccount(w HttpWriter, r HttpReq, info map[string]any) any {
	acc, ok := GetUid(w, r, false)
	if acc == nil {acc = new(Account)}
	return map[string]any{
		"loggedin": ok,
		"accountId"  : bog(acc==nil, 0, acc.AccountId),
		"accountName": bog(acc==nil, "<NOT LOGGED IN>", acc.AccountName),
		"isAdmin"    : bog(acc==nil, false, acc.IsAdmin),
	}
}

