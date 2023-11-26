<template>
  <q-list class="shici-layout">
    <q-item v-if="dynasties.length > 0">
      <q-item-section side top>
        <q-item-label>
          <q-btn flat :color="!form.dynasty?'primary':'black'" label="朝代" @click="handleRemoveQuery('dynasty')"></q-btn>
        </q-item-label>
      </q-item-section>
      <q-item-section>
        <q-item-label>
          <template v-for="(dynasty, index) in dynasties" :key="index">
            <q-btn flat
                   :label="dynasty.name"
                   :color="form.dynasty == dynasty.id ?'primary':'black'"
                   @click="handleQuery({dynasty: dynasty.id})" />
          </template>
        </q-item-label>
      </q-item-section>
    </q-item>

    <q-item>
      <q-item-section>
        <q-item-label>
          <q-btn flat icon="sort" label="默认排序" @click="handleQuery({sort: 'desc'})" v-if="!form.sort" />
          <q-btn flat icon="sort" color="primary" label="倒序" @click="handleQuery({sort: 'random'})" v-else-if="form.sort == 'desc'" />
          <q-btn flat icon="sort" color="primary" label="随机" @click="handleRemoveQuery('sort')" v-else />
        </q-item-label>
      </q-item-section>
      <q-space />
      <q-item-section side>
        <q-item-label>
          <q-input dense outlined label="输入作者名称" @keyup.enter="handleAuthors" v-model="result.form.name">
            <template v-slot:append>
              <q-icon name="search" style="cursor: pointer;" @click="handleAuthors" />
            </template>
          </q-input>
        </q-item-label>
      </q-item-section>
    </q-item>

    <q-item v-if="result.list.length == 0">
      <q-item-section>
        <q-item-label class="flex flex-center text-grey">无数据</q-item-label>
      </q-item-section>
    </q-item>
    <template v-else>
      <q-item>
        <q-item-section>
          <div class="row q-gutter-sm full-width">
            <q-card flat bordered :key="index" v-for="(row, index) in result.list" class="col" style="min-width: 32%;">
              <q-card-section class="row justify-between">
                <div class="text-h6 ellipsis self-baseline" style="cursor: pointer; max-width: 80%" @click="handleAuthor(row)">
                  {{ row.name }}
                </div>
                <div class="text-subtitle2 ellipsis text-primary self-baseline" v-if="row.dynasty">{{ row.dynasty.name }}</div>
              </q-card-section>

              <template v-if="row.desc != ''">
                <q-separator inset />

                <q-card-section>
                  <div class="ellipsis-3-lines" style="white-space: pre-line;">
                    {{ row.desc }}
                  </div>
                </q-card-section>
              </template>
            </q-card>
          </div>
        </q-item-section>
      </q-item>

      <q-item v-if="paginationTotal > 1">
        <q-item-section class="items-center">
          <q-pagination direction-links boundary-numbers
                        size="sm"
                        :max="paginationTotal"
                        :max-pages="5"
                        :model-value="result.pagination.page"
                        @update:model-value="handlePagination" />
        </q-item-section>
      </q-item>
    </template>
    <q-inner-loading label="加载中" :showing="result.loading"></q-inner-loading>
  </q-list>
</template>

<script setup>
 import { computed, ref, watch, onMounted } from 'vue';
 import { getCurrentInstance } from 'vue'

 const { proxy } = getCurrentInstance()

 const dynasties = ref([])

 const form = computed(() => {
     return proxy.$route.query
 })

 const result = ref({
     list: [],
     form: {},
     loading: false,
     pagination: {
         page: 1,
         rowsPerPage: 0,
         rowsNumber: 0,
     }
 })

 const paginationTotal = computed(() => {
     if (result.value.list.length == result.value.pagination.rowsNumber) {
         return 1
     }
     return Math.ceil(result.value.pagination.rowsNumber / result.value.pagination.rowsPerPage)
 })

 const deepClone = (data) => {
     return JSON.parse(JSON.stringify(data));
 }

 const handlePagination = (page) => {
     result.value.pagination.page = page
     handleAuthors()
 }

 const handleQuery = (query) => {
     proxy.$router.push({path: "/authors", query: {...form.value, ...query}})
 }

 const handleRemoveQuery = (key) => {
     let query = {...form.value}
     delete query[key]
     proxy.$router.push({path: "/authors", query: query})
 }

 const handleAuthor = (row) => {
     proxy.$router.push({path: `/authors/${row.id}`})
 }

 const handleAuthors = () => {
     const params = {...form.value}
     if (result.value.form.name) {
         params.name = result.value.form.name
     }
     if (result.value.pagination.page > 1) {
         params.page = result.value.pagination.page
     }

     result.value.loading = true
     proxy.$api.get("/api/authors", {
         params: params
     }).then(resp => {
         result.value.pagination.rowsNumber = resp.data.data.total
         result.value.pagination.rowsPerPage = resp.data.data.limit
         result.value.list = resp.data.data.list
     }).finally(_ => {
         result.value.loading = false
     })
 }

 const handleDynasties = () => {
     proxy.$api.get("/api/dynasties").then(resp => {
         dynasties.value = resp.data.data.list
     })
 }

 watch(() => form.value,
       (newQuery, oldQuery) => {
           result.value.pagination.page = 1
           handleAuthors()
       }
 )

 onMounted(() => {
     handleAuthors()
     handleDynasties()
 })
</script>
