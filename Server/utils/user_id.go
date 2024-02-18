package utils

import "sync"

var (
    userID string
    lock   sync.Mutex
    isSetUserId  bool
)
// ユーザーIDをセットする
func SetUserID(id string) {
    lock.Lock()
    defer lock.Unlock()
    
    // すでに設定されている場合は何もしない
    if isSetUserId {
        return
    }

    userID = id
		// UserIdを設定済みにする
    isSetUserId = true
}

// ユーザーIDを取得する
func GetUserID() string {
    lock.Lock()
    defer lock.Unlock()
    return userID
}

