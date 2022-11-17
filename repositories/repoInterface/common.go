package repoInterface

type RepoInterface interface {
	SetOrm(orm any) //该方法主要是在事务中 修改当前仓库的session
	ResetOrm()      //该方法用户还原被修改的orm
}
