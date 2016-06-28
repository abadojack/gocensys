package gocensys

//PrimarySeries represents primary series of the scan.
type PrimarySeries struct {
	AllX509Certificates       PrimarySeriesDetail `json:"All X.509 Certificates"`
	IPv4Snapshots             PrimarySeriesDetail `json:"IPv4 Snapshots"`
	AlexaTop1MillionSnapshots PrimarySeriesDetail `json:"Alexa Top 1 Million Snapshots"`
}

type PrimarySeriesDetail struct {
	Description  string       `json:"description"`
	DetailsURL   string       `json:"details_url"`
	LatestResult LatestResult `json:"latest_result"`
	ID           string       `json:"id"`
	Name         string       `json:"name"`
}

type LatestResult struct {
	Timestamp  string `json:"timestamp"`
	Name       string `json:"name"`
	DetailsURL string `json:"details_url"`
}

type RawSeries struct {
	Two2SSHBannerFullIpv4             RawSeriesDetail `json:"22-ssh-banner-full_ipv4"`
	Two5SMTPStarttlsAlexaTop1Mil      RawSeriesDetail `json:"25-smtp-starttls-alexa_top1mil"`
	One10Pop3StarttlsFullIpv4         RawSeriesDetail `json:"110-pop3-starttls-full_ipv4"`
	Four43HTTPSSsl3FullIpv4           RawSeriesDetail `json:"443-https-ssl_3-full_ipv4"`
	Four43HTTPSHeartbleedFullIpv4     RawSeriesDetail `json:"443-https-heartbleed-full_ipv4"`
	One911FoxDeviceIDFullIpv4         RawSeriesDetail `json:"1911-fox-device_id-full_ipv4"`
	Four43HTTPSDheFullIpv4            RawSeriesDetail `json:"443-https-dhe-full_ipv4"`
	ZeroIcmpEchoRequestFullIpv4       RawSeriesDetail `json:"0-icmp-echo_request-full_ipv4"`
	Four65SmtpsSsl2FullIpv4           RawSeriesDetail `json:"465-smtps-ssl_2-full_ipv4"`
	Nine95Pop3SRsaExportFullIpv4      RawSeriesDetail `json:"995-pop3s-rsa_export-full_ipv4"`
	Two0000Dnp3StatusFullIpv4         RawSeriesDetail `json:"20000-dnp3-status-full_ipv4"`
	One02S7SzlFullIpv4                RawSeriesDetail `json:"102-s7-szl-full_ipv4"`
	Four43HTTPSHeartbleedAlexaTop1Mil RawSeriesDetail `json:"443-https-heartbleed-alexa_top1mil"`
	Two5SMTPDheExport1SampleIpv4      RawSeriesDetail `json:"25-smtp-dhe_export-1%_sample_ipv4"`
	Two5SMTPSsl2FullIpv4              RawSeriesDetail `json:"25-smtp-ssl_2-full_ipv4"`
	Four43HTTPSDheAlexaTop1Mil        RawSeriesDetail `json:"443-https-dhe-alexa_top1mil"`
	Eight0HTTPGetFullIpv4             RawSeriesDetail `json:"80-http-get-full_ipv4"`
	Nine93ImapsSsl2FullIpv4           RawSeriesDetail `json:"993-imaps-ssl_2-full_ipv4"`
	Nine93ImapsTLSFullIpv4            RawSeriesDetail `json:"993-imaps-tls-full_ipv4"`
	Two3TelnetBannerFullIpv4          RawSeriesDetail `json:"23-telnet-banner-full_ipv4"`
	Four43HTTPSDheExportAlexaTop1Mil  RawSeriesDetail `json:"443-https-dhe_export-alexa_top1mil"`
	Two5SMTPRsaExportFullIpv4         RawSeriesDetail `json:"25-smtp-rsa_export-full_ipv4"`
	Eight0HTTPOpenProxyFullIpv4       RawSeriesDetail `json:"80-http-open_proxy-full_ipv4"`
	Eight0HTTPGetAlexaTop1Mil         RawSeriesDetail `json:"80-http-get-alexa_top1mil"`
	One43ImapSsl2FullIpv4             RawSeriesDetail `json:"143-imap-ssl_2-full_ipv4"`
	Four43HTTPSSsl2FullIpv4           RawSeriesDetail `json:"443-https-ssl_2-full_ipv4"`
	Nine93ImapsRsaExportFullIpv4      RawSeriesDetail `json:"993-imaps-rsa_export-full_ipv4"`
	One10Pop3Ssl2FullIpv4             RawSeriesDetail `json:"110-pop3-ssl_2-full_ipv4"`
	Four43HTTPSRsaExportAlexaTop1Mil  RawSeriesDetail `json:"443-https-rsa_export-alexa_top1mil"`
	Four43HTTPSTLSFullIpv4            RawSeriesDetail `json:"443-https-tls-full_ipv4"`
	Four7808BacnetDeviceIDFullIpv4    RawSeriesDetail `json:"47808-bacnet-device_id-full_ipv4"`
	One43ImapStarttlsFullIpv4         RawSeriesDetail `json:"143-imap-starttls-full_ipv4"`
	Five3DNSLookupFullIpv4            RawSeriesDetail `json:"53-dns-lookup-full_ipv4"`
	Four43HTTPSRsaExport1SampleIpv4   RawSeriesDetail `json:"443-https-rsa_export-1%_sample_ipv4"`
	Seven547CwmpGetFullIpv4           RawSeriesDetail `json:"7547-cwmp-get-full_ipv4"`
	Two5SMTPStarttlsFullIpv4          RawSeriesDetail `json:"25-smtp-starttls-full_ipv4"`
	Two1FtpBannerFullIpv4             RawSeriesDetail `json:"21-ftp-banner-full_ipv4"`
	Four43HTTPSDheExport1SampleIpv4   RawSeriesDetail `json:"443-https-dhe_export-1%_sample_ipv4"`
	Four65SmtpsTLSFullIpv4            RawSeriesDetail `json:"465-smtps-tls-full_ipv4"`
	Five02ModbusDeviceIDFullIpv4      RawSeriesDetail `json:"502-modbus-device_id-full_ipv4"`
	Nine95Pop3STLSFullIpv4            RawSeriesDetail `json:"995-pop3s-tls-full_ipv4"`
	Four43HTTPSRsaExportFullIpv4      RawSeriesDetail `json:"443-https-rsa_export-full_ipv4"`
	Four43HTTPSSsl3AlexaTop1Mil       RawSeriesDetail `json:"443-https-ssl_3-alexa_top1mil"`
	Four43HTTPSDheExportFullIpv4      RawSeriesDetail `json:"443-https-dhe_export-full_ipv4"`
	Four43HTTPSHeartbleed1SampleIpv4  RawSeriesDetail `json:"443-https-heartbleed-1%_sample_ipv4"`
	Nine93ImapsDheExportFullIpv4      RawSeriesDetail `json:"993-imaps-dhe_export-full_ipv4"`
	Four43HTTPSTLSAlexaTop1Mil        RawSeriesDetail `json:"443-https-tls-alexa_top1mil"`
	Nine95Pop3SSsl2FullIpv4           RawSeriesDetail `json:"995-pop3s-ssl_2-full_ipv4"`
}

type RawSeriesDetail struct {
	Subprotocol  string       `json:"subprotocol"`
	Desription   string       `json:"desription"`
	Protocol     string       `json:"protocol"`
	Name         string       `json:"name"`
	DetailsURL   string       `json:"details_url"`
	LatestResult LatestResult `json:"latest_result"`
	Destination  string       `json:"destination"`
	ID           string       `json:"id"`
	Port         int          `json:"port"`
}

type RawSeriesDetails struct {
	Subprotocol  string       `json:"subprotocol"`
	Desription   string       `json:"desription"`
	Protocol     string       `json:"protocol"`
	Name         string       `json:"name"`
	DetailsURL   string       `json:"details_url"`
	LatestResult LatestResult `json:"latest_result"`
	Destination  string       `json:"destination"`
	ID           string       `json:"id"`
	Port         int          `json:"port"`
}

type Result struct {
	HistoricalResult []HistoricalResult `json:"historical"`
	LatestResult     LatestResult       `json:"latest_result"`
}

type HistoricalResult struct {
	Timestamp  string `json:"timestamp"`
	Name       string `json:"name"`
	DetailsURL string `json:"details_url"`
}

type ScanResult struct {
	Files  Files       `json:"files"`
	TaskID interface{} `json:"task_id"`
	Series struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"series"`
	PrimaryFile File   `json:"primary_file"`
	Timestamp   string `json:"timestamp"`
	ID          string `json:"id"`
	Metadata    struct {
	} `json:"metadata"`
}

type Files struct {
	ZtagMetadata     File `json:"ztag-metadata"`
	ZteeZgrabUpdates File `json:"ztee-zgrab-updates"`
	ZgrabMetadata    File `json:"zgrab-metadata"`
	ZgrabResults     File `json:"zgrab-results"`
	ZmapLog          File `json:"zmap-log"`
	ZgrabLog         File `json:"zgrab-log"`
	ZteeZmapUpdates  File `json:"ztee-zmap-updates"`
	ZmapResults      File `json:"zmap-results"`
	ZmapMetadata     File `json:"zmap-metadata"`
	ZtagLog          File `json:"ztag-log"`
}

type File struct {
	CompressedSize              int         `json:"compressed_size"`
	CompressedSha256Fingerprint interface{} `json:"compressed_sha256_fingerprint"`
	CompressedDownloadPath      string      `json:"compressed_download_path"`
	Sha256Fingerprint           string      `json:"sha256_fingerprint"`
	DownloadPath                string      `json:"download_path"`
	FileType                    string      `json:"file_type"`
	Schema                      interface{} `json:"schema"`
	CompressionType             string      `json:"compression_type"`
	Size                        int         `json:"size"`
}
