package data

import ()

// GetConfig Gets the config
func GetConfig() (Config, error) {
	var config Config
	keyValues := make(map[string]string)

	rows, err := Database.Query(`SELECT *
                               FROM Config`)

	if err != nil {
		return config, err
	}

	for rows.Next() {
		var key string
		var value string
		err = rows.Scan(&key, &value)

		if err != nil {
			return config, err
		}

		keyValues[key] = value
	}

	config.KeyValues = keyValues

	return config, nil
}

// CreateConfig Create the config
func CreateConfig(config Config) (Config, error) {
	_, err := Database.Exec(`INSERT INTO Config
                              ()
                           VALUES
                              ()`)
	return config, err
}

// UpdateConfig Updates the config
func UpdateConfig(config Config) (Config, error) {

	return config, nil
}
