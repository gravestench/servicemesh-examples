package static_assets

func (m *Middleware) DependenciesResolved() bool {
	if m.router == nil {
		return false
	}

	return true
}

func (m *Middleware) ResolveDependencies(mesh servicemesh.M) {
	for _, candidate := range rt.Services() {
		if router, ok := candidate.(IsWebRouter); ok {
			if router.RouteRoot() == nil {
				return
			}

			m.router = router
		}
	}
}
