package models



func GetAppId() []byte {
	randData := RandomData()
	if randData == nil {
		return nil
	}
	return GetMd5(randData)
}

func GetAppKey() []byte {
	randData := RandomData()
	if randData == nil {
		return nil
	}
	return GetMd5(randData)
}
