package common

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/btcsuite/btcutil/base58"
)

type UID struct {
	localID    uint64
	objectType uint16
	shardID    uint16
}

func NewUID(localId uint64, objectType, shardId uint16) UID {
	return UID{
		localID:    localId,
		objectType: objectType,
		shardID:    shardId,
	}
}

func (uid UID) String() string {
	val := uid.localID<<28 | uint64(uid.objectType)<<18 | uint64(uid.shardID)<<0
	return base58.Encode([]byte(fmt.Sprintf("%v", val)))
}

func (uid UID) GetLocalID() uint64 {
	return uid.localID
}

func (uid UID) GetObjectType() uint16 {
	return uid.objectType
}

func (uid UID) GetShardID() uint16 {
	return uid.shardID
}

func DecomposeUID(s string) (UID, error) {
	uid, err := strconv.ParseUint(s, 10, 64)

	if err != nil {
		return UID{}, err
	}

	if (1 << 18) > uid {
		return UID{}, ErrWrongUID
	}

	u := UID{
		localID:    uint64(uid >> 28),
		objectType: uint16(uid >> 18 & 0x3FF),
		shardID:    uint16(uid >> 0 & 0x3FFFF),
	}

	return u, nil
}

func DecomposeUIDFromBase58(s string) (UID, error) {
	return DecomposeUID(string(base58.Decode(s)))
}

// Parse from fake_id or id to id
func ParseId(id interface{}) (uint64, error) {

	_, isStr := id.(string)
	_, isNum := id.(uint64)
	if !isStr && !isNum {
		return 0, ErrInvalidRequest(nil)
	}
	if isStr {
		uid, err := DecomposeUIDFromBase58(id.(string))
		if err != nil {
			return 0, ErrInvalidRequest(err)
		}
		return uid.GetLocalID(), nil
	} else {
		return id.(uint64), nil
	}
}

func (uid UID) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", uid.String())), nil
}

func (uid *UID) UnmarshalJSON(data []byte) error {
	decodeUID, err := DecomposeUIDFromBase58(strings.Replace(string(data), "\"", "", -1))

	if err != nil {
		return err
	}

	uid.localID = decodeUID.localID
	uid.objectType = decodeUID.objectType
	uid.shardID = decodeUID.shardID

	return nil
}

func (uid *UID) Value() (driver.Value, error) {
	if uid == nil {
		return nil, nil
	}

	return int64(uid.localID), nil
}

func (uid *UID) Scan(value interface{}, objectId, shardId *uint16) error {
	if value == nil {
		return nil
	}

	var i uint64

	switch t := value.(type) {
	case int:
		i = uint64(t)
	case int8:
		i = uint64(t)
	case int16:
		i = uint64(t)
	case int32:
		i = uint64(t)
	case int64:
		i = uint64(t)
	case uint8:
		i = uint64(t)
	case uint16:
		i = uint64(t)
	case uint32:
		i = uint64(t)
	case uint64:
		i = uint64(t)
	case []byte:
		a, err := strconv.Atoi(string(t))
		if err != nil {
			return err
		}

		i = uint64(a)
	case string:
		id, err := DecomposeUIDFromBase58(t)
		if err != nil {
			return err
		}
		*uid = id
		return nil
	default:
		return errors.New("invalid Scan Source")
	}
	var (
		_objectId uint16
		_shardId  uint16
	)
	if objectId != nil {
		_objectId = *objectId
	}
	if shardId != nil {
		_shardId = *shardId
	}
	*uid = NewUID(i, _objectId, _shardId)
	return nil
}
