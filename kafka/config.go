package kafka

const (
	// TopicNameUserNames receives user names that should be scraped.
	TopicNameUserNames = "user_names"
	// TopicNameUserInfos receives scraped user info
	TopicNameUserInfos = "user_infos"
	// TopicNameUserScrapeErrors contains alls scrape errors
	TopicNameUserScrapeErrors = "user_scrape_errors"
	// TopicNameScrapedPosts receives scraped posts.
	// Queued for insertion to the database.
	TopicNameScrapedPosts = "scraped_posts"
	// TopicNamePostScrapeErrors receives errors that happen in scraping
	TopicNamePostScrapeErrors = "post_scrape_errors"
	// TopicNameScrapedLikes receives scraped likes.
	// Queued for insertion to the database.
	TopicNameScrapedLikes = "scraped_likes"
	// TopicNameLikeScrapeErrors receives errors that happen in scraping
	TopicNameLikeScrapeErrors = "like_scrape_errors"
	// TopicNameScrapedComments receives scraped comments.
	// Queued for insertion to the database.
	TopicNameScrapedComments = "scraped_comments"
	// TopicNameCommentScrapeErrors receives errors that happen in scraping
	TopicNameCommentScrapeErrors = "comment_scrape_errors"

	// TopicNamePGPosts is a compacted topic for the posts that are written to postgres
	TopicNamePGPosts = "postgres.public.posts"

	// TopicNamePGUsers is a compacted topic for the users that are written to postgres
	TopicNamePGUsers = "postgres.public.users"

	// TopicNamePGFaces is a compacted topic for the users that are written to postgres
	TopicNamePGFaces = "postgres.public.faces"

	// TopicNamePictureDownloads receives messages for the picture downloads
	TopicNamePictureDownloads = "download_jobs"

	// TopicNameReconJobs receives messages for each newly downloaded post
	TopicNameReconJobs = "recon_jobs"

	// TopicNameReconResults receives a message for analysed post that contains faces
	TopicNameReconResults = "recon_results"
)

// WithVersion returns the topic or groupid with a version suffix
// func WithVersion(topicName string, version int) string {
// 	return fmt.Sprintf("%s_v%d", topicName, version)
// }
