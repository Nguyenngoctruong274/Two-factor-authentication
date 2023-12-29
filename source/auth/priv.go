package auth

// var serCfg map[string]string

// func initJWTKey() (err error) {
// 	ctx := context.Background()
// 	privKeyName := getConfig(ctx, "jwtPrivKey")
// 	passWord := "{}{}{}"
// 	if signKey, err = jwt.ParseRSAPrivateKeyFromPEMWithPassword([]byte(privKeyName), passWord); err != nil {
// 		return err
// 	}
// 	// if verifyKey, err = jwt.ParseRSAPublicKeyFromPEM([]byte(privKeyName)); err != nil {
// 	// 	return err
// 	// }
// 	return
// }

// func getConfig(ctx context.Context, key string) string {
// 	yamlFile, err := ioutil.ReadFile("/home/wee/ProjectOwn/authentication/config.yaml")
// 	if err != nil {
// 		log.Fatalf("Failed to read YAML file: %v", err)
// 	}

// 	if err := yaml.Unmarshal(yamlFile, &serCfg); err != nil {
// 		log.Fatalf("Failed to unmarshal YAML: %v", err)
// 	}
// 	return serCfg[key]
// }
