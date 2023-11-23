<template>
  <q-card flat>
    <q-card-section class="column justify-center items-center">
      <div class="text-h6">
        {{ row.name }}
      </div>
      <div class="row q-gutter-sm text-primary">
        <span class="text-subtitle2">{{ row.dynasty.name }}</span>
      </div>
    </q-card-section>

    <q-card-section class="column justify-center items-center" v-if="row.desc != ''">
      <div style="white-space: pre-line; font-size: 1.05rem;">
        {{ row.desc }}
      </div>
    </q-card-section>

    <q-card-section v-if="row.short_desc != ''">
      <span class="text-subtitle2 text-primary">简介：</span>
      <div style="white-space: pre-line;">
        {{ row.short_desc }}
      </div>
    </q-card-section>

    <q-card-section>
      <span class="text-subtitle2 text-primary">相关诗词：</span>
      <div class="row q-gutter-sm">
        <q-card flat bordered :key="index" v-for="(poem, index) in row.poems" style="width: 32%;">
          <template v-if="poem.title != ''">
            <q-card-section class="row justify-between">
              <div class="text-h6 ellipsis self-baseline" style="cursor: pointer; max-width: 80%;"
                   @click="handlePoem(poem)">
                {{ poem.title }}<template v-if="poem.chapter != ''"> · {{ poem.chapter }}</template>
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
  </q-card>
</template>

<script setup>
 import { computed, ref, watch, onMounted } from 'vue';
 import { getCurrentInstance } from 'vue'

 const { proxy } = getCurrentInstance()

 const row = ref({
     dynasty:{},
 })

 const handleAuthor = () => {
     proxy.$api.get(`/api/authors/${proxy.$route.params.id}`).then(resp => {
         row.value = resp.data.data

         proxy.$api.get(`/api/poems`, {
             params: {author: row.value.id}
         }).then(resp => {
             row.value.poems = resp.data.data.list
         })
     })
 }

 const handlePoem = (row) => {
     proxy.$router.push({path: `/poems/${row.id}`})
 }

 onMounted(() => {
     handleAuthor()
 })
</script>