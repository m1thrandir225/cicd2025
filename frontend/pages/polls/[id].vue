<template>
  <div v-if="status === 'pending'">Loading...</div>
  <div
    v-else-if="status === 'success' && data"
    class="mx-auto my-4 flex max-w-[650px] flex-col items-start gap-8 transition-all duration-300 ease-in-out"
  >
    <div class="flex w-full flex-row items-start justify-between">
      <h1>{{ data.poll.description }}</h1>
      <div class="flex flex-row items-start gap-4">
        <p class="">
          Active until <br />
          <span class="text-[12px] text-orange-600">{{
            formatDate(new Date(data.poll.active_until))
          }}</span>
        </p>
        <button
          class="flex items-center justify-center rounded-md border border-neutral-400 px-4 py-2 text-neutral-900 transition-all duration-300 ease-in-out hover:border-transparent hover:bg-neutral-200 hover:text-orange-600 dark:text-neutral-100 dark:hover:bg-neutral-800"
          @click="toggleQr()"
        >
          <Icon name="material-symbols:share" size="18px" mode="css" />
        </button>
      </div>
    </div>
    <Transition name="fade">
      <PollShare v-if="qrShown" :url="route.fullPath" />
    </Transition>
    <PollVote
      v-if="!hasVoted && !isPollByUser"
      :options="data.options"
      :refresh-page-data="refresh"
    />
    <div v-else class="flex w-full flex-col items-start gap-8">
      <h1>Poll Stats:</h1>
      <PollStatsChart v-if="chartData" :chart-data="chartData" />
      <PollStatsTable :data="data.options" />
    </div>
  </div>
</template>

<script setup lang="ts">
import PollShare from "~/components/polls/PollShare.vue";
import PollStatsChart from "~/components/polls/PollStatsChart.vue";
import PollStatsTable from "~/components/polls/PollStatsTable.vue";
import PollVote from "~/components/polls/PollVote.vue";
const route = useRoute();

const qrShown = ref(false);

const toggleQr = useToggle(qrShown);

const { data, error, status, refresh } = await useFetch(
  `/api/polls/${route.params.id}`,
  {
    method: "GET",
  },
);
useSeoMeta({
  title: data ? `${data.value?.poll.description}` : "poll details",
});

const voteStore = useVoteStore();

const authStore = useAuthStore();

const isPollByUser = computed(() => {
  return data.value?.poll.created_by === authStore.user?.id;
});

const hasVoted = computed(() => {
  if (!data.value?.options) return false;

  return data.value.options.some((el) => {
    const contained = voteStore.findVoteByOptionId(el.id);
    // If contained is an array, check if it has any elements
    return Array.isArray(contained) ? contained.length > 0 : !!contained;
  });
});

const chartData = computed(() => {
  if (!data.value || !data.value?.options) return null;

  return buildChartData(data.value?.options);
});

if (error.value) {
  throw createError({
    statusCode: error.value.statusCode,
    fatal: true,
  });
}
</script>
