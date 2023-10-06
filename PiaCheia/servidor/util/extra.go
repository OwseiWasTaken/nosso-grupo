package util

import (
	"hash/fnv"
	"os"
	"errors"
	"net/http"
	"strings"
	"fmt"
	"strconv"
)

type HttpWriter = http.ResponseWriter
type HttpReq = *http.Request

func Exists(filepath string) (bool) {
	_, err := os.Stat(filepath);
	return !errors.Is(err, os.ErrNotExist)
}

type HashResult = uint32
const HashBitLen = 32
func Hash(s string) HashResult {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()+uint32(90749*len(s))
}

func Panic(e error) {
	if (e != nil) {
		panic(e)
	}
}

func Unpack[T any](v T, e error) T {
	if (e != nil) {
		panic(e)
	}
	return v
}

func ReadBasicAuth(str string) (name string, value HashResult, ok bool) {
	steps := strings.SplitN(str, ":", 2)
	if len(steps) < 2 { // no ":" in string
		return "", 0 ,false
	}
	value64, ValidHashResult := strconv.ParseUint(steps[1], 10, HashBitLen)
	if (ValidHashResult != nil) {
		return "", 0 ,false
	}
	return steps[0], HashResult(value64), true
}

func MakeBasicAuth(name string, value HashResult) string {
	return fmt.Sprintf("%s:%v", name, value)
}

func ServingError(w HttpWriter, r HttpReq, err string, code int) {
	fmt.Fprintf(w, "ERRO CÓDIGO %d (%s): %s", code, http.StatusText(code), err)
	//w.WriteHeader(code)
}

func Redirect(
	w HttpWriter, r HttpReq,
	url string,
	reason string, args... interface{},
) {
	//Log(USER, "%s -> %s: %s\n", r.URL.Path, url, fmt.Sprintf(reason, args...))
	http.Redirect(w, r, url, http.StatusSeeOther)
}

const UidCookieName = "piacheia-uid"
func GetUid(w HttpWriter, r HttpReq, mustLogin bool) (acc *Account, loggedInt bool) {
	ReadCookie, err := r.Cookie(UidCookieName)
	if (errors.Is(err, http.ErrNoCookie)) {
		if (mustLogin) {
			Redirect(w, r, "/login", "user without login cookie")
		}
		return nil, false
	}

	if (err != nil) {
		//Log(GO+USER+ERROR, "failed to read cookie: %s\n", err)
		if (mustLogin) {
			ServingError(w, r, "Sintaxe de cookie inválida", 400)
		}
		return nil, false
	}

	uid, Passhash, ValidCookieAuth := ReadBasicAuth(ReadCookie.Value)
	if (!ValidCookieAuth) {
		//Log(GO+USER+ERROR, "failed to read cookie BasicAuth: %s\n", ReadCookie.Value)
		if (mustLogin) {
			ServingError(w, r, "Sintaxe de valores dentro de cookie", 400)
		}
		return nil, false
	}

	//accountsLock.Lock()
	//defer accountsLock.Unlock()
	acc, ok := NameToAccount[uid]
	if (!ok || acc.Passhash != Passhash) {
		if (mustLogin) {
			Redirect(w, r, "/login", "Login inválido")
		}
		return nil, false
	}
	return acc, true
}

func SetUid(w HttpWriter, name string, passhash HashResult) {
	SendCookie := http.Cookie{
		Name: UidCookieName,
		Value: MakeBasicAuth(name, passhash),
		HttpOnly: true,
		//Secure: true, // not serving https yet
		Path: "/",
		SameSite: http.SameSiteStrictMode,
		Expires: time.Now().Add(time.Hour * 24 * 10), // 10 dias
	}
	http.SetCookie(w, &SendCookie)
}

func bog[A any](expr bool, iftrue, iffalse A) (A) {
	if (expr) {
		return iftrue
	}
	return iffalse
}
