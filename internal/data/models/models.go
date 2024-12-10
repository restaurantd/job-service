package models

import "job-service/internal/proto/pb"

type Worker struct {
	Firstname  string
	Secondname string
	Advtitle   string
}

type Adv struct {
	Salary      int
	Title       string
	Description string
}

func ToPB(p []Adv) []*pb.Adv {
	var advs []*pb.Adv
	for _, a := range p {
		advs = append(advs,
			&pb.Adv{
				Title:       a.Title,
				Description: a.Description,
				Salary:      int64(a.Salary),
			},
		)
	}

	return advs
}

func ToPbWorkers(p []Worker) []*pb.Worker {
	var workers []*pb.Worker

	for _, s := range p {
		workers = append(workers,
			&pb.Worker{
				Firstname:  s.Firstname,
				Secondname: s.Secondname,
				AdvTitle:   s.Advtitle,
			},
		)
	}

	return workers
}
