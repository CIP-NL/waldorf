// Package Waldorf implements a simple checking framework for doing multiple assertions and collecting the
// result of the assertions. Waldorf does not use interfaces but instead relies on the user providing statements
// which evaluate to booleans, making it typed as opposed to the assertions library by stretchr.
package waldorf

func (r *Observation) ShouldBeTrue(statement bool, msg string, formatting ...interface{}) (validated bool) {
	if !statement {
		r.complaintsMu.Lock()
		defer r.complaintsMu.Unlock()
		r.complaints.comps = append(r.complaints.comps, generateMSG(msg, formatting...))
	}
	return statement
}

func (r *Observation) ShouldBeFalse(statement bool, msg string, formatting ...interface{}) (validated bool) {
	if statement {
		r.complaintsMu.Lock()
		defer r.complaintsMu.Unlock()
		r.complaints.comps = append(r.complaints.comps, generateMSG(msg, formatting...))
	}
	return !statement
}

func (r *Observation) EitherShouldBeTrue(statementOne, statementTwo bool, msg string, formatting ...interface{}) (validated bool) {
	if !(statementOne || statementTwo) {
		r.complaintsMu.Lock()
		defer r.complaintsMu.Unlock()
		r.complaints.comps = append(r.complaints.comps, generateMSG(msg, formatting...))
	}
	return statementOne || statementTwo
}

func (r *Observation) BothShouldBeTrue(statementOne, statementTwo bool, msg string, formatting ...interface{}) (validated bool) {
	if !(statementOne && statementTwo) {
		r.complaintsMu.Lock()
		defer r.complaintsMu.Unlock()
		r.complaints.comps = append(r.complaints.comps, generateMSG(msg, formatting...))
	}
	return statementOne && statementTwo
}

func (r *Observation) BothShouldBeFalse(statementOne, statementTwo bool, msg string, formatting ...interface{}) (validated bool) {
	if statementOne || statementTwo {
		r.complaintsMu.Lock()
		defer r.complaintsMu.Unlock()
		r.complaints.comps = append(r.complaints.comps, generateMSG(msg, formatting...))
	}
	return !(statementOne || statementTwo)
}

func (r *Observation) EitherNeither(statementOne, statementTwo bool, msg string, formatting ...interface{}) (validated bool) {
	if statementOne && statementTwo {
		r.complaintsMu.Lock()
		defer r.complaintsMu.Unlock()
		r.complaints.comps = append(r.complaints.comps, generateMSG(msg, formatting...))
	}
	return !(statementOne && statementTwo)
}

func (r *Observation) IfThisThenThat(statementOne, statementTwo bool, msg string, formatting ...interface{}) (validated bool) {
	if !statementOne {
		return true
	}

	if !statementTwo {
		r.complaintsMu.Lock()
		defer r.complaintsMu.Unlock()
		r.complaints.comps = append(r.complaints.comps, generateMSG(msg, formatting...))
		return false
	}
	return true
}
