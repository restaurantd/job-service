package server

import (
	"context"
	"job-service/internal/data"
	"job-service/internal/data/models"
	"job-service/internal/loger"
	"job-service/internal/proto/pb"
)

type Server struct {
	Loger loger.Logger
	DB    data.DataBase
	pb.UnimplementedJobServer
}

func (s Server) GetWorkers(_ context.Context, req *pb.GetWorkersReq) (*pb.GetWorkersResp, error) {
	workers, err := s.DB.GetWorkers(req.Title)
	if err != nil {
		return &pb.GetWorkersResp{}, err
	}

	ws := models.ToPbWorkers(workers)

	return &pb.GetWorkersResp{Workers: ws}, nil
}
func (s Server) AddWorker(_ context.Context, req *pb.AddWorkerReq) (*pb.Resp, error) {
	w := models.Worker{
		Firstname:  req.Worker.Firstname,
		Secondname: req.Worker.Secondname,
		Advtitle:   req.Worker.AdvTitle,
	}

	if err := s.DB.Add(w); err != nil {
		return &pb.Resp{Status: err.Error()}, err
	}

	return &pb.Resp{
		Status: "succes",
	}, nil
}
func (s Server) GetAdv(_ context.Context, req *pb.ReqGetAdv) (*pb.Adv, error) {
	adv, err := s.DB.GetAdv(req.Title)
	if err != nil {
		return &pb.Adv{}, err
	}

	return &pb.Adv{
		Title:       adv.Title,
		Description: adv.Description,
		Salary:      int64(adv.Salary),
	}, nil
}
