package bean

type Url struct{
	UUID  string `dao:"u_uuid"`
	Title string `form:"title" binding:"required" dao:"u_title"`
	Url   string `form:"url" binding:"required" dao:"u_url"`
	Date  string `form:"date" binding:"required" dao:"u_date"`
	Desc  string `form:"desc" binding:"required" dao:"u_desc"`
}




