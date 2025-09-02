package core

type UPID string

type QEMU struct {
	Id   uint64 `json:"id"`
	Node uint64 `json:"node"`
}

func (q *QEMU) Create() (*UPID, error) {
	return nil, nil
}

func (q *QEMU) Delete() (*UPID, error) {
	return nil, nil
}

func (q *QEMU) GetConfig() (*UPID, error) {
	return nil, nil
}

func (q *QEMU) UpdateConfig() (*UPID, error) {
	return nil, nil
}
