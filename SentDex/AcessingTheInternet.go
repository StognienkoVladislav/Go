package main

import (
	"fmt"
	"encoding/xml"
)

var washPostXML = []byte(`
<sitemapindex>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-politics-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-blogs-politics-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-opinions-sitemap.xml</loc>
   </sitemap>
</sitemapindex>
`)

type Location struct{
	Loc string `xml:"loc"`
}


type Sitemapindex struct{
	Locations []Location `xml:"sitemap"`
}

func (e Location) String () string{
	return fmt.Sprint(e.Loc)
}

func main()  {
	bytes:= washPostXML
	var s Sitemapindex
	xml.Unmarshal(bytes, &s)
	fmt.Println(s.Locations)
	}
