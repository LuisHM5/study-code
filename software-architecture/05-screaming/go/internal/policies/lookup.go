package policies

// Find consulta una póliza por su id.
func (d *Desk) Find(id string) (Policy, error) {
	p, ok := d.store.byIDOrNil(id)
	if !ok {
		return Policy{}, ErrNotFound
	}
	return p, nil
}
