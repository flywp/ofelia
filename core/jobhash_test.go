package core

import (
	. "gopkg.in/check.v1"
)

type SuiteJobHash struct{}

var _ = Suite(&SuiteJobHash{})

func newHashExecJob(container, user string) *ExecJob {
	return &ExecJob{
		BareJob:   BareJob{Name: "wpcron", Schedule: "@every 10m", Command: "wp cron event run --due-now"},
		Container: container,
		User:      user,
	}
}

func (s *SuiteJobHash) TestExecJobHashDetectsFieldChanges(c *C) {
	base := newHashExecJob("site-php-1", "www-data")

	user := newHashExecJob("site-php-1", "")
	c.Assert(user.Hash(), Not(Equals), base.Hash())

	container := newHashExecJob("other-php-1", "www-data")
	c.Assert(container.Hash(), Not(Equals), base.Hash())

	tty := newHashExecJob("site-php-1", "www-data")
	tty.TTY = true
	c.Assert(tty.Hash(), Not(Equals), base.Hash())

	env := newHashExecJob("site-php-1", "www-data")
	env.Environment = []string{"FOO=bar"}
	c.Assert(env.Hash(), Not(Equals), base.Hash())
}

func (s *SuiteJobHash) TestExecJobHashIgnoresClientAndRuntimeState(c *C) {
	a := newHashExecJob("site-php-1", "1000:1000")
	b := newHashExecJob("site-php-1", "1000:1000")
	b.Client = &mockDockerClient{}
	b.cronID = 42
	b.execID = "abc"
	b.history = append(b.history, &Execution{})
	c.Assert(b.Hash(), Equals, a.Hash())
}

func (s *SuiteJobHash) TestRunJobHashDetectsFieldChanges(c *C) {
	newJob := func() *RunJob {
		return &RunJob{BareJob: BareJob{Name: "r", Schedule: "@hourly", Command: "true"}, Image: "alpine", User: "nobody"}
	}
	base := newJob()

	user := newJob()
	user.User = "1000"
	c.Assert(user.Hash(), Not(Equals), base.Hash())

	image := newJob()
	image.Image = "busybox"
	c.Assert(image.Hash(), Not(Equals), base.Hash())

	withClient := newJob()
	withClient.Client = &mockDockerClient{}
	c.Assert(withClient.Hash(), Equals, base.Hash())
}

func (s *SuiteJobHash) TestRunServiceJobHashDetectsFieldChanges(c *C) {
	newJob := func() *RunServiceJob {
		return &RunServiceJob{BareJob: BareJob{Name: "s", Schedule: "@hourly", Command: "true"}, Image: "alpine"}
	}
	base := newJob()

	user := newJob()
	user.User = "1000"
	c.Assert(user.Hash(), Not(Equals), base.Hash())

	withClient := newJob()
	withClient.Client = &mockDockerClient{}
	c.Assert(withClient.Hash(), Equals, base.Hash())
}

func (s *SuiteJobHash) TestLocalJobHashDetectsFieldChanges(c *C) {
	base := &LocalJob{BareJob: BareJob{Name: "l", Schedule: "@hourly", Command: "true"}}
	dir := &LocalJob{BareJob: BareJob{Name: "l", Schedule: "@hourly", Command: "true"}, Dir: "/tmp"}
	c.Assert(dir.Hash(), Not(Equals), base.Hash())
}
