package xar

import "encoding/xml"

type xar struct {
	XMLName xml.Name `xml:"xar"`
	Toc     toc
}

type toc struct {
	XMLName      xml.Name `xml:"toc"`
	CreationTime string   `xml:"creation-time"`
	Checksum     checksum
	Files        []file `xml:"file"`
}

type checksum struct {
	XMLName xml.Name
	Style   string `xml:"style,attr"`
	Offset  string `xml:"offset"`
	Size    string `xml:"size"`
}

type signature struct {
	XMLName xml.Name
	Style   string `xml:"style,attr"`
	Offset  string `xml:"offset"`
	Size    string `xml:"size"`
}

type keyInfo struct {
	XMLName  xml.Name
	X509Data x509Data `xml:"X509Data"`
}

type x509Data struct {
	XMLName         xml.Name `xml:"X509Data"`
	X509Certificate []string `xml:"X509Certificate"`
}

type file struct {
	XMLName          xml.Name `xml:"file"`
	Id               string   `xml:"id,attr"`
	Data             fileData
	FinderCreateTime fileFinderCreateTime
	Ctime            string `xml:"ctime"`
	Mtime            string `xml:"mtime"`
	Atime            string `xml:"atime"`
	Group            string `xml:"group"`
	Gid              int    `xml:"gid"`
	User             string `xml:"user"`
	Uid              int    `xml:"uid"`
	Mode             uint32 `xml:"mode"`
	DeviceNo         uint64 `xml:"deviceno"`
	Inode            uint64 `xml:"inode"`
	Type             string `xml:"type"`
	Name             string `xml:"name"`
	Files            []file `xml:"file"`
}

type fileData struct {
	XMLName           xml.Name `xml:"data"`
	Len               uint64   `xml:"length"`
	Encoding          fileEncoding
	Offset            uint64       `xml:"offset"`
	Size              uint64       `xml:"size"`
	ExtractedChecksum fileChecksum `xml:"extracted-checksum"`
	ArchivedChecksum  fileChecksum `xml:"archived-checksum"`
}

type fileEncoding struct {
	XMLName xml.Name `xml:"encoding"`
	Style   string   `xml:"style,attr"`
}

type fileChecksum struct {
	XMLName xml.Name
	Style   string `xml:"style,attr"`
	Digest  string `xml:",chardata"`
}

type fileFinderCreateTime struct {
	XMLName     xml.Name `xml:"FinderCreateTime"`
	Nanoseconds int64    `xml:"nanoseconds"`
	Time        string   `xml:"time"`
}
