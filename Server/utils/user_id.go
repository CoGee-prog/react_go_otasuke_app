package utils

import "sync"

var (
    userID string
    lock   sync.Mutex
    isSet  bool
)
// ユーザーIDをセットする
func SetUserID(id string) {
    lock.Lock()
    defer lock.Unlock()
    
    // すでに設定されている場合は何もしない
    if isSet {
        return
    }

    userID = id
		// UserIdを設定済みにする
    isSet = true
}

// ユーザーIDを取得する
func GetUserID() string {
    lock.Lock()
    defer lock.Unlock()
    return userID
}

