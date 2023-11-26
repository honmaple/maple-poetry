<template>
  <q-card flat class="shici-layout">
    <q-card-section class="column flex-center">
      <div class="text-h6">
        <template v-if="result.row.chapter != ''">{{ result.row.chapter }} · </template>{{ result.row.title }}
      </div>
      <div class="row q-gutter-sm text-primary">
        <span class="text-subtitle2">{{ result.row.dynasty.name }}</span>
        <span class="text-subtitle2">{{ result.row.author.name }}</span>
      </div>
    </q-card-section>

    <q-card-section class="column flex-center">
      <div class="shici-content">
        {{ result.row.content }}
      </div>
    </q-card-section>

    <q-card-section v-if="!result.loading && result.row.author && result.row.author.desc != ''">
      <span class="text-subtitle2 text-primary">作者：</span>
      <div style="white-space: pre-line;">
        {{ result.row.author.desc }}
      </div>
    </q-card-section>

    <q-card-section v-if="!result.loading && result.row.annotation != ''">
      <span class="text-subtitle2 text-primary">注解：</span>
      <div style="white-space: pre-line;">
        {{ result.row.annotation }}
      </div>
    </q-card-section>
    <q-card-section class="row" :class="{'justify-end': !result.row.prev, 'justify-start': !result.row.next, 'justify-between': result.row.prev && result.row.next}">
      <q-btn flat class="primary text-primary" icon="chevron_left"
             @click="$router.push({params: {id: result.row.prev.id}})" v-if="result.row.prev">
        <template v-if="result.row.prev.chapter != ''">{{ result.row.prev.chapter }} · </template>
        {{ result.row.prev.title }}
      </q-btn>
      <q-btn flat class="primary text-primary self-end" icon-right="chevron_right"
             @click="$router.push({params: {id: result.row.next.id}})" v-if="result.row.next">
        <template v-if="result.row.next.chapter != ''">{{ result.row.next.chapter }} · </template>
        {{ result.row.next.title }}
      </q-btn>
    </q-card-section>
    <q-inner-loading label="加载中" :showing="result.loading"></q-inner-loading>
  </q-card>
</template>

<script setup>
 import { computed, ref, watch, onMounted } from 'vue';
 import { getCurrentInstance } from 'vue'

 const { proxy } = getCurrentInstance()

 const result = ref({
     row: {
         author:{},
         dynasty:{},
     },
     loading: false,
 })

 const handlePoem = () => {
     result.value.loading = true

     proxy.$api.get(`/api/poems/${proxy.$route.params.id}`).then(resp => {
         result.value.row = resp.data.data
     }).finally(_ => {
         result.value.loading = false
     })
 }

 onMounted(() => {
     handlePoem()
 })
</script>
