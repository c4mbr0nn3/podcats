<template>
  <v-row>
    <v-col cols="1" class="d-flex align-center justify-center">
      <v-icon icon="mdi-close" size="small" @click="setRead(notification.ID)" />
    </v-col>
    <v-col cols="11">
      <v-card variant="tonal" @click="onClick">
        <v-card-text>
          <v-row>
            <v-col cols="2">
              <v-avatar :image="podcast.Image" />
            </v-col>
            <v-col cols="10">
              {{ message }}
              <div class="text-right text-caption">{{ timestamp }}</div>
            </v-col>
          </v-row>
        </v-card-text>
      </v-card>
    </v-col>
  </v-row>
</template>

<script setup>
import { computed } from "vue";
import { usePodcastsStore } from "@/stores/podcasts";
import { useNotificationsStore } from "@/stores/notifications";
import { useRouter } from "vue-router";
import { formatDate } from "@/utils/date";

const props = defineProps({
  notification: {
    type: Object,
    required: true,
  },
});

const podcastsStore = usePodcastsStore();
const notificationsStore = useNotificationsStore();
const { setRead } = notificationsStore;

const router = useRouter();

const podcast = computed(() => {
  const { PodcastId } = props.notification;
  const podcast = podcastsStore.podcasts.find((p) => p.ID === PodcastId);
  return podcast;
});

const timestamp = computed(() => {
  const { PublicationDate } = props.notification;
  return formatDate(PublicationDate);
});

const message = computed(() => {
  if (props.notification.Message.length < 70) return props.notification.Message;
  return props.notification.Message.substring(0, 70) + "...";
});

const onClick = () => {
  const { ID, PodcastId, PodcastItemId } = props.notification;
  notificationsStore.setRead(ID);
  router.push({
    name: "single-item",
    params: { id: PodcastId, itemId: PodcastItemId },
  });
};
</script>
