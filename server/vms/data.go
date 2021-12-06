package vms

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Balancer struct {
	Id int64 `json:"id"`
	Shutdown []int64 `json:"usedVMS"`
	All int64 `json:"allVMS"`
}

type Container struct {
	Db *pgxpool.Pool
}

func CreateContainer(db *pgxpool.Pool) *Container {
	return &Container{Db: db}
}

func (cont *Container) ConnInfo() ([]*Balancer, error){

	balancers, fault := cont.Db.Query(context.Background(), "select distinct b.id, count(*) over(partition by b.id) from balancers b join vms on b.id = vms.balancerid order by b.id;")
	
	if fault != nil {
		return nil, fault
	}

	defer balancers.Close()

	res := make([]*Balancer, 0)

	stdVMS, fault := cont.getStdVMS()

	if fault != nil {
		return nil, fault
	}

	for balancers.Next() {

		var balancerid int64
		var amount int64

		if fault := balancers.Scan(&balancerid, &amount); fault != nil {
			return nil, fault
		}

		std := stdVMS[balancerid]
		if std == nil {
			std = make([]int64, 0)
		}

		el := &Balancer{Id: balancerid, Shutdown: std, All: amount}
		res = append(res, el)

	}

	return res, nil

}

func (cont *Container) getStdVMS() (map[int64][]int64, error) {

	stdVMS, err1 := cont.Db.Query(context.Background(), "select balancerid, id from vms where status = true")

	stdVMSarr := map[int64][]int64{}

	if err1 != nil {
		return nil, err1
	}

	defer stdVMS.Close()

	for stdVMS.Next() {

		var balancerid int64
		var vmsid int64

		if fault := stdVMS.Scan(&balancerid, &vmsid); fault != nil {
			return nil, fault
		}

		stdVMSarr[balancerid] = append(stdVMSarr[balancerid], vmsid)

	}

	return stdVMSarr, nil
}

func (cont *Container) UpdateVMS(vmsid int64, state bool) error {

	_, fault := cont.Db.Exec(context.Background(),"update vms set status = $1 where id = $2", state, vmsid)

	return fault

}
