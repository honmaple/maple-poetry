<template>
  <div class="full-height full-width">
    <q-carousel transition-prev="scale" transition-next="scale"
                swipeable
                infinite
                animated
                control-color="white"
                padding
                arrows
                class="bg-primary text-white full-height full-width"
                v-model="slide">
      <q-carousel-slide :name="index" class="column no-wrap flex-center" :key="index" v-for="(row, index) in result.list">
        <q-card flat class="bg-primary text-white">
          <q-card-section class="column flex-center">
            <div class="text-h6">
              <template v-if="row.chapter != ''">{{ row.chapter }} · </template>{{ row.title }}
            </div>
            <div class="row q-gutter-sm">
              <span class="text-subtitle2" v-if="row.dynasty">{{ row.dynasty.name }}</span>
              <span class="text-subtitle2" v-if="row.author">{{ row.author.name }}</span>
            </div>
          </q-card-section>
          <q-card-section class="column flex-center">
            <div class="shici-content">
              {{ row.content }}
            </div>
          </q-card-section>
        </q-card>
      </q-carousel-slide>
    </q-carousel>
    <q-inner-loading class="bg-primary text-white" style="z-index:999;" label="诗词加载中" :showing="result.loading"></q-inner-loading>
  </div>
</template>

<script setup>
 import { ref, onMounted } from 'vue';
 import { getCurrentInstance } from 'vue'

 const { proxy } = getCurrentInstance()

 const slide = ref(0)

 const result = ref({
     list: [],
     loading: false,
 })

 const handlePoem = () => {
     result.value.loading = true

     proxy.$api.get(`/api/poems`, {
         params: {sort: "random"}
     }).then(resp => {
         result.value.list = resp.data.data.list
     }).finally(_ => {
         result.value.loading = false
     })
 }

 onMounted(() => {
     handlePoem()
 })
</script>