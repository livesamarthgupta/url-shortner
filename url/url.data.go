package url

import (
	"github.com/lib/pq"
	"github.com/urlshortner/database"
)

func SetupTable() error {
	_, err := database.DBConn.Exec(`CREATE SCHEMA IF NOT EXISTS urlshortner;`)
	if err != nil {
		return err
	}

	_, err = database.DBConn.Exec(`CREATE TABLE IF NOT EXISTS urlshortner.urls
	(originalurl VARCHAR(255) NOT NULL,
	shortenurl VARCHAR(255) NOT NULL,
	PRIMARY KEY(shortenurl));`)
	if err != nil {
		return err
	}

	return nil
}

func insertOrUpdateUrl(url URL) error {
	_, err := database.DBConn.Exec(`INSERT INTO urlshortner.urls
	(
		originalurl,
		shortenurl
	)
	VALUES 
	(
		$1, $2
	);`,
		url.OriginalUrl,
		url.ShortenUrl,
	)
	if err != nil {
		pqErr := err.(*pq.Error)
		if(pqErr.Code == pq.ErrorCode("23505")) {
			updateUrl(url)
		} else {
			return err
		}
	}
	return nil
}

func updateUrl(url URL) error {
	_, err := database.DBConn.Exec(`UPDATE urlshortner.urls SET originalurl = $1 WHERE shortenurl = $2;`,
		url.OriginalUrl,
		url.ShortenUrl,
	)
	if err != nil {
		return err
	}
	return nil
}

func fetchOriginalUrl(shortenUrl string) (string, error) {
	result := database.DBConn.QueryRow(`SELECT originalurl from urlshortner.urls WHERE shortenurl = $1;`, shortenUrl)
	var originalUrl string
	err := result.Scan(&originalUrl)
	if err != nil {
		return "", err
	}

	return originalUrl, nil
}
