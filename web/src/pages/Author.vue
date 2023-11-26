<template>
  <q-card flat class="shici-layout">
    <q-card-section class="column justify-center items-center">
      <div class="text-h6">
        {{ result.row.name }}
      </div>
      <div class="row q-gutter-sm text-primary">
        <span class="text-subtitle2">{{ result.row.dynasty.name }}</span>
      </div>
    </q-card-section>

    <q-card-section class="column justify-center items-center" v-if="result.row.desc != ''">
      <div style="white-space: pre-line; font-size: 1.05rem;">
        {{ result.row.desc }}
      </div>
    </q-card-section>

    <q-card-section v-if="!result.loading && result.row.short_desc != ''">
      <span class="text-subtitle2 text-primary">简介：</span>
      <div style="white-space: pre-line;">
        {{ result.row.short_desc }}
      </div>
    </q-card-section>

    <q-card-section v-if="result.row.poems && result.row.poems.length > 0">
      <span class="text-subtitle2 text-primary">相关诗词：</span>
      <div class="row q-gutter-sm">
        <q-card flat bordered :key="index" v-for="(poem, index) in result.row.poems" style="width: 32%;">
          <template v-if="poem.title != ''">
            <q-card-section class="row justify-between">
              <div class="text-h6 ellipsis self-baseline" style="cursor: pointer; max-width: 80%;"
                   @click="handlePoem(poem)">
                {{ poem.title }}<template v-if="poem.chapter != ''"> · {{ poem.chapter }}</template>
              </div>
              <div class="text-subtitle2 ellipsis text-primary self-baseline" v-if="result.row.dynasty">
                {{ result.row.dynasty.name }}
              </div>
            </q-card-section>
            <q-separator inset />
          </template>

          <q-card-section>
            <div class="ellipsis-3-lines" style="white-space: pre-line;">
              {{ poem.content }}
            </div>
          </q-card-section>
        </q-card>
      </div>
    </q-card-section>

    <q-card-section class="row" :class="{'justify-end': !result.row.prev, 'justify-start': !result.row.next, 'justify-between': result.row.prev && result.row.next}">
      <q-btn flat class="primary text-primary" icon="chevron_left"
             @click="$router.push({params: {id: result.row.prev.id}})" v-if="result.row.prev">
        {{ result.row.prev.name }}
      </q-btn>
      <q-btn flat class="primary text-primary self-end" icon-right="chevron_right"
             @click="$router.push({params: {id: result.row.next.id}})" v-if="result.row.next">
        {{ result.row.next.name }}
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
     row: ref({
         dynasty:{},
     }),
     loading: false,
 })

 const handleAuthor = () => {
     result.value.loading = true
     proxy.$api.get(`/api/authors/${proxy.$route.params.id}`).then(resp => {
         result.value.row = resp.data.data

         proxy.$api.get(`/api/poems`, {
             params: {author: result.value.row.id}
         }).then(resp => {
             result.value.row.poems = resp.data.data.list
         })
     }).finally(_ => {
         result.value.loading = false
     })
 }

 const handlePoem = (row) => {
     proxy.$router.push({path: `/poems/${row.id}`})
 }

 onMounted(() => {
     handleAuthor()
 })
</script>