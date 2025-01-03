using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;

namespace Infrastructure.Entities;

[Table("episodes")]
public class Episode : AuditableEntity
{
    [Column("id")] [Key] public Guid Id { get; set; }
    [Column("title")] [Required] public string Title { get; set; }
    [Column("summary")] [Required] public string Summary { get; set; }
    [Column("episode_tye")] [Required] public string EpisodeType { get; set; }
    [Column("published_date")] [Required] public DateTime PublishedDate { get; set; }
    [Column("file_url")] [Required] public string FileUrl { get; set; }
    [Column("guid")] [Required] public string Guid { get; set; }
    [Column("image")] [Required] public string Image { get; set; }

    [Column("last_play_position")]
    [Required]
    public int LastPlayPosition { get; set; }

    [Column("is_played")] [Required] public bool IsPlayed { get; set; }
    [Column("bookmark_date")] [Required] public DateTime BookmarkDate { get; set; }
    [Column("podcast_id")] [Required] public Guid PodcastId { get; set; }
    [Column("user_id")] [Required] public Guid UserId { get; set; }

    public Podcast Podcast { get; set; }
    public User User { get; set; }
}