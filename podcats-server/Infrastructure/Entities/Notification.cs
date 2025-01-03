using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;

namespace Infrastructure.Entities;

[Table("notifications")]
public class Notification : AuditableEntity
{
    [Column("id")] [Key] public Guid Id { get; set; }
    [Column("user_id")] [Required] public Guid UserId { get; set; }
    [Column("podcast_id")] [Required] public Guid PodcastId { get; set; }
    [Column("episode_id")] [Required] public Guid EpisodeId { get; set; }
    [Column("published_date")] [Required] public DateTime PublishedDate { get; set; }
    [Column("is_read")] [Required] public bool IsRead { get; set; }
    [Column("message")] [Required] public string Message { get; set; }

    public User User { get; set; }
    public Podcast Podcast { get; set; }
    public Episode Episode { get; set; }
}