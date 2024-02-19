package jobque

type jobResult struct {
	err    error
	result interface{}
}

type jobQueFn struct {
	fn       func() (result interface{}, err error)
	resultCh chan *jobResult
}

type JobQues struct {
	queCh chan *jobQueFn
}

func (self *JobQues) loop() {
	for job := range self.queCh {
		go func(job *jobQueFn) {
			res, err := job.fn()
			job.resultCh <- &jobResult{
				err:    err,
				result: res,
			}
		}(job)
	}
}

func (self *JobQues) Exec(fn func() (interface{}, error)) (interface{}, error) {
	job := &jobQueFn{
		fn:       fn,
		resultCh: make(chan *jobResult),
	}

	go func() {
		self.queCh <- job
	}()

	result := <-job.resultCh

	return result.result, result.err
}

func NewJobQues() *JobQues {
	q := &JobQues{
		queCh: make(chan *jobQueFn),
	}

	go q.loop()

	return q
}
