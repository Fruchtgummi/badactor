package badactor

import (
	"strconv"
	"testing"
	"time"
)

func TestStudioStartReaper(t *testing.T) {

	//var b bool
	var err error

	st := NewStudio(256)
	an := "actorname"
	rn := "rulename"
	rm := "rulemessage"

	r := &Rule{
		Name:        rn,
		Message:     rm,
		StrikeLimit: 3,
		ExpireBase:  time.Second * 10,
		Sentence:    time.Minute * 10,
	}

	// add rule
	st.AddRule(r)

	// creat directors
	err = st.CreateDirectors(1024)
	if err != nil {
		t.Errorf("CreateDirectors failed %v %v", an, rn)
	}

	dur := time.Millisecond * time.Duration(1)
	st.StartReaper(dur)

	msg := st.Status()
	if !msg.reaperAlive {
		t.Errorf("isAlive should be true %v", msg.reaperAlive)
	}

	msg = st.Status()
	if !msg.reaperAlive {
		t.Errorf("isAlive should be true %v", msg.reaperAlive)
	}

}

func TestStudioIsJailedFor(t *testing.T) {
	var b bool
	var err error

	st := NewStudio(256)
	an := "actorname"
	rn := "rulename"
	rm := "rulemessage"
	r := &Rule{
		Name:        rn,
		Message:     rm,
		StrikeLimit: 3,
		ExpireBase:  time.Second * 10,
		Sentence:    time.Minute * 10,
	}

	// add rule
	st.AddRule(r)

	// creat directors
	err = st.CreateDirectors(1024)
	if err != nil {
		t.Errorf("CreateDirectors failed %v %v", an, rn)
	}

	b = st.IsJailedFor(an, rn)
	if b {
		t.Errorf("IsJailed should be false, %v %v", an, rn)
	}

	for i := 0; i < 3; i++ {
		err = st.Infraction(an, rn)
		if err != nil {
			t.Errorf("Infraction failed %v, %v", an, rn)
		}
	}

	b = st.IsJailedFor(an, rn)
	if !b {
		t.Errorf("IsJailed should be true, %v %v", an, rn)
	}

}

func TestStudioIsJailed(t *testing.T) {
	var b bool
	var err error

	st := NewStudio(256)
	an := "actorname"
	rn := "rulename"
	rm := "rulemessage"
	r := &Rule{
		Name:        rn,
		Message:     rm,
		StrikeLimit: 3,
		ExpireBase:  time.Second * 10,
		Sentence:    time.Minute * 10,
	}

	// add rule
	st.AddRule(r)

	// creat directors
	err = st.CreateDirectors(1024)
	if err != nil {
		t.Errorf("CreateDirectors failed %v %v", an, rn)
	}

	b = st.IsJailed(an)
	if b {
		t.Errorf("IsJailed should be false, %v %v", an, rn)
	}

	for i := 0; i < 3; i++ {
		err = st.Infraction(an, rn)
		if err != nil {
			t.Errorf("Infraction failed %v, %v", an, rn)
		}
	}

	b = st.IsJailed(an)
	if !b {
		t.Errorf("IsJailed should be true, %v %v", an, rn)
	}

}

func TestStudioKeepAlive(t *testing.T) {
	var ti time.Time
	var err error

	st := NewStudio(256)
	an := "actorname"
	rn := "rulename"
	rm := "rulemessage"
	r := &Rule{
		Name:        rn,
		Message:     rm,
		StrikeLimit: 3,
		ExpireBase:  time.Second * 10,
		Sentence:    time.Minute * 10,
	}

	// add rule
	st.AddRule(r)

	// creat directors
	err = st.CreateDirectors(1024)
	if err != nil {
		t.Errorf("CreateDirectors failed %v %v", an, rn)
	}

	d := st.Director(an)
	ti, err = d.lTimeToLive(an)
	if err == nil {
		t.Errorf("lTimeToLive() should fail as actor doesn't exist %v, %v, %v", an, err, ti)
	}

	err = st.Infraction(an, rn)
	if err != nil {
		t.Errorf("Infraction failed %v, %v", an, rn)
	}

	orig_ttl, err := d.lTimeToLive(an)
	if err != nil {
		t.Errorf("lTimeToLive() should not fail as actor exists %v, %v, %v", an, err, orig_ttl)
	}

	err = st.KeepAlive(an)
	if err != nil {
		t.Errorf("KeepAlive failed %v, %v", an, rn)
	}

	new_ttl, err := d.lTimeToLive(an)
	if err != nil {
		t.Errorf("lTimeToLive() should not fail as actor exists %v, %v, %v", an, err, orig_ttl)
	}

	if new_ttl == orig_ttl {
		t.Errorf("KeepAlive() failed to update time %v, %v, %v", an, orig_ttl, new_ttl)
	}

}

func TestStudioCreateActor(t *testing.T) {
	var err error

	st := NewStudio(256)
	an := "actorname"
	rn := "rulename"
	rm := "rulemessage"
	r := &Rule{
		Name:        rn,
		Message:     rm,
		StrikeLimit: 3,
		ExpireBase:  time.Second * 10,
		Sentence:    time.Minute * 10,
	}

	// add rule
	st.AddRule(r)

	// creat directors
	err = st.CreateDirectors(1024)
	if err != nil {
		t.Errorf("CreateDirectors failed %v %v", an, rn)
	}

	if st.ActorExists(an) {
		t.Errorf("ActorExists should be false")
	}

	err = st.CreateActor(an, rn)
	if err != nil {
		t.Errorf("CreateActor should be nil [%v]", err)
	}

	if !st.ActorExists(an) {
		t.Errorf("InfractionExists should be true")
	}
}

func TestStudioCreateInfraction(t *testing.T) {
	var err error

	st := NewStudio(256)
	an := "actorname"
	rn := "rulename"
	rm := "rulemessage"
	r := &Rule{
		Name:        rn,
		Message:     rm,
		StrikeLimit: 3,
		ExpireBase:  time.Second * 10,
		Sentence:    time.Minute * 10,
	}

	// add rule
	st.AddRule(r)

	// creat directors
	err = st.CreateDirectors(1024)
	if err != nil {
		t.Errorf("CreateDirectors failed %v %v", an, rn)
	}

	if st.InfractionExists(an, rn) {
		t.Errorf("InfractionExists should be false")
	}

	err = st.CreateActor(an, rn)
	if err != nil {
		t.Errorf("CreateActor should be nil [%v]", err)
	}

	err = st.CreateInfraction(an, rn)
	if err != nil {
		t.Errorf("CreateInfraction should be nil [%v]", err)
	}

	if !st.InfractionExists(an, rn) {
		t.Errorf("InfractionExists should be true")
	}
}

func TestStudioStrikes(t *testing.T) {
	var si int
	var err error

	st := NewStudio(256)
	an := "actorname"
	rn := "rulename"
	rm := "rulemessage"
	r := &Rule{
		Name:        rn,
		Message:     rm,
		StrikeLimit: 3,
		ExpireBase:  time.Second * 10,
		Sentence:    time.Minute * 10,
	}

	// add rule
	st.AddRule(r)

	// creat directors
	err = st.CreateDirectors(1024)
	if err != nil {
		t.Errorf("CreateDirectors failed %v %v", an, rn)
	}

	si, err = st.Strikes(an, rn)
	if si != 0 {
		t.Errorf("Strikes for Actor [%v] and Rule [%v] should not be %v, %v", an, rn, si, err)
	}

	// 1st inf
	err = st.Infraction(an, rn)
	if err != nil {
		t.Errorf("Infraction failed for Actor [%v] and Rule [%v] should not be %v", an, rn, err)
	}

	si, err = st.Strikes(an, rn)
	if si != 1 {
		t.Errorf("Strikes for Actor [%v] and Rule [%v] should not be %v %v", an, rn, si, err)
	}

	// 2nd inf
	st.Infraction(an, rn)
	si, err = st.Strikes(an, rn)
	if si != 2 {
		t.Errorf("Strikes for Actor [%v] and Rule [%v] should not be %v %v", an, rn, si, err)
	}

	// 3rd inf, jail, strikes for that infraction name should be 0
	st.Infraction(an, rn)
	si, err = st.Strikes(an, rn)
	if si != 0 {
		t.Errorf("Strikes for Actor [%v] and Rule [%v] should not be %v %v", an, rn, si, err)
	}

	// should still be jailed
	st.Infraction(an, rn)
	si, err = st.Strikes(an, rn)
	if si != 0 {
		t.Errorf("Strikes for Actor [%v] and Rule [%v] should not be %v %v", an, rn, si, err)
	}
}

func TestStudioAddRule(t *testing.T) {
	st := NewStudio(256)
	an := "an_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	rn := "rn_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	rm := "rm_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	r := &Rule{
		Name:        rn,
		Message:     rm,
		StrikeLimit: 3,
		ExpireBase:  time.Millisecond * 10,
		Sentence:    time.Millisecond * 10,
	}

	// add rule
	st.AddRule(r)

	// test rule exists
	_, ok := st.rules[r.Name]
	if ok == false {
		t.Errorf("AddRule for Actor [%s] should not fail", an)
	}
}

func TestStudioAddRules(t *testing.T) {
	st := NewStudio(2)
	r1 := &Rule{
		Name:        "rule1",
		Message:     "message1",
		StrikeLimit: 3,
		ExpireBase:  time.Millisecond * 10,
		Sentence:    time.Millisecond * 10,
	}

	r2 := &Rule{
		Name:        "rule2",
		Message:     "message2",
		StrikeLimit: 3,
		ExpireBase:  time.Millisecond * 10,
		Sentence:    time.Millisecond * 10,
	}

	// add rule safety is of no concern
	st.AddRule(r1)
	st.AddRule(r2)

	// apply rules
	err := st.CreateDirectors(1024)
	if err != nil {
		t.Errorf("CreateDirectors failed %v", err)
	}

	// range of rules and directors and make sure the rule exists for each
	for di, d := range st.directors {
		if !d.ruleExists(r1.Name) {
			t.Errorf("ApplyRules for director [%v] is missing rule %v", di, r1.Name)
		}
		if !d.ruleExists(r2.Name) {
			t.Errorf("ApplyRules for director [%v] is missing rule %v", di, r2.Name)
		}
	}

}

func TestStudioCreateDirectors(t *testing.T) {

	var nocap int32
	nocap = 256

	st := NewStudio(256)
	rn := "rn_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	rm := "rm_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	r := &Rule{
		Name:        rn,
		Message:     rm,
		StrikeLimit: 3,
		ExpireBase:  time.Millisecond * 10,
		Sentence:    time.Millisecond * 10,
	}

	// add rule safety is of no concern
	st.AddRule(r)

	if st.capacity != nocap {
		t.Errorf("Capacity for Studio want [%v] got [%v]", st.capacity, nocap)
	}

	st.CreateDirectors(256)

	var i int32
	for i = 0; i < st.capacity; i++ {
		_, ok := st.directors[i]
		if !ok {
			t.Errorf("Director [%v] for Studio was not created", i)
		}
	}

}
