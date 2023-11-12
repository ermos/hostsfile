package hostsfile

type HostOption func(host *Host)

// NewHost returns a new host.
func NewHost(address string, hostNames []string, options ...HostOption) *Host {
	host := &Host{
		address:   address,
		hostNames: hostNames,
	}

	for _, opt := range options {
		opt(host)
	}

	return host
}

// WithComment returns a host option to set the comment of the host.
func WithComment(comment string) HostOption {
	return func(host *Host) {
		host.comment = comment
	}
}
