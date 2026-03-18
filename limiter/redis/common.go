package redis

func generateKey(domain, id string) string {
	return domain + "/id:" + id

}
