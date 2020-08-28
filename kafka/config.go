package kafka

const (
	// TopicNameUserNames receives user names that should be scraped.
	TopicNameUserNames = "user_names"

	// TopicNameScrapedPosts receives scraped posts.
	// Queued for insertion to the database.
	TopicNameScrapedPosts = "scraped_posts"
	// TopicNamePostScrapeErrors receives errors that happen in scraping
	TopicNamePostScrapeErrors = "post_scrape_errors"
)

// WithVersion returns the topic or groupid with a version suffix
// func WithVersion(topicName string, version int) string {
// 	return fmt.Sprintf("%s_v%d", topicName, version)
// }
