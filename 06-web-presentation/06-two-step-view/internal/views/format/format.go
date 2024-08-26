package format

import (
	"fmt"

	"github.com/jibaru/two-step-view/internal/views/logical"
)

const (
	Head = `<!DOCTYPE html>
			<html lang="en">
			  <head>
			    <meta charset="UTF-8" />
			    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
			    <title>Album Details</title>
			    <style>
			      body {
			        font-family: Arial, sans-serif;
			        background-color: #f4f4f4;
			        margin: 0;
			        padding: 20px;
			      }
			
			      h1 {
			        color: #333;
			        text-align: center;
			      }
			
			      p {
			        background-color: #fff;
			        padding: 15px;
			        border-radius: 5px;
			        box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
			        max-width: 600px;
			        margin: 10px auto;
			        color: #555;
			      }
			    </style>
			  </head>
			  <body>`
	Foot = `</body>
			</html>`
)

func RenderAlbumToHTML(album logical.RenderedAlbum) string {
	return fmt.Sprintf("<h1>%s</h1><p>Artist: %s</p>", album.Title, album.Artist)
}
