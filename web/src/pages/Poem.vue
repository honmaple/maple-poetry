<template>
  <q-card flat>
    <q-card-section class="column justify-center items-center">
      <div class="text-h6">
        {{ row.title }}<template v-if="row.chapter != ''"> · {{ row.chapter }}</template>
      </div>
      <div class="row q-gutter-sm text-primary">
        <span class="text-subtitle2">{{ row.dynasty.name }}</span>
        <span class="text-subtitle2">{{ row.author.name }}</span>
      </div>
    </q-card-section>

    <q-card-section class="column justify-center items-center">
      <div style="white-space: pre-line; font-size: 1.05rem;">
        {{ row.content }}
      </div>
    </q-card-section>

    <q-card-section v-if="row.author && row.author.desc != ''">
      <span class="text-subtitle2 text-primary">作者：</span>
      <div style="white-space: pre-line;">
        {{ row.author.desc }}
      </div>
    </q-card-section>

    <q-card-section v-if="row.note != ''">
      <span class="text-subtitle2 text-primary">注解：</span>
      <div style="white-space: pre-line;">
        {{ row.note }}
      </div>
    </q-card-section>
  </q-card>
</template>

<script setup>
 import { computed, ref, watch, onMounted } from 'vue';
 import { getCurrentInstance } from 'vue'

 const { proxy } = getCurrentInstance()

 const row = ref({
     author:{},
     dynasty:{},
 })

 const handlePoem = () => {
     proxy.$api.get(`/api/poems/${proxy.$route.params.id}`).then(resp => {
         row.value = resp.data.data
     })
 }

 onMounted(() => {
     handlePoem()
 })
</script>
