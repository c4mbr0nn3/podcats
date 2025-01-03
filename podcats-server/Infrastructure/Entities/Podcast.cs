using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;

namespace Infrastructure.Entities;

[Table("podcasts")]
public class Podcast : AuditableEntity
{
    [Column("id")] [Key] public Guid Id { get; set; }
    [Column("title")] [Required] public string Title { get; set; }
    [Column("summary")] [Required] public string Summary { get; set; }
    [Column("author")] [Required] public string Author { get; set; }
    [Column("image")] [Required] public string Image { get; set; }
    [Column("show_url")] [Required] public string ShowUrl { get; set; }
    [Column("feed_url")] [Required] public string FeedUrl { get; set; }
    [Column("episodes_count")] [Required] public int EpisodesCount { get; set; }
    [Column("played_count")] [Required] public int PlayedCount { get; set; }
    [Column("latest_episode")] [Required] public DateTime LatestEpisode { get; set; }
    [Column("user_id")] [Required] public Guid UserId { get; set; }

    public User User { get; set; }
    public List<Episode> Episodes { get; set; }
}