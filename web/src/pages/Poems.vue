<template>
  <q-list>
    <q-item v-if="dynastyResult.list.length > 0">
      <q-item-section side top>
        <q-item-label>
          <q-btn flat :color="!form.dynasty?'primary':'black'" label="朝代" @click="handleRemoveQuery('dynasty')"></q-btn>
        </q-item-label>
      </q-item-section>
      <q-item-section>
        <q-item-label>
          <template v-for="(dynasty, index) in dynastyResult.list" :key="index">
            <q-btn flat
                   :label="dynasty.name"
                   :color="form.dynasty == dynasty.id ?'primary':'black'"
                   @click="handleQuery({dynasty: dynasty.id})" />
          </template>
        </q-item-label>
      </q-item-section>
    </q-item>

    <q-item v-if="collectionResult.list.length > 0">
      <q-item-section side top>
        <q-item-label>
          <q-btn flat :color="!form.collection?'primary':'black'" label="诗集" @click="handleRemoveQuery('collection')"></q-btn>
        </q-item-label>
      </q-item-section>
      <q-item-section>
        <q-item-label>
          <template v-for="(collection, index) in collectionResult.list" :key="index">
            <q-btn flat
                   :label="collection.name"
                   :color="form.collection == collection.id ?'primary':'black'"
                   @click="handleQuery({collection: collection.id})" />
          </template>
          <q-btn flat label="更多" to="/collections" v-if="collectionResult.pagination.rowsNumber > collectionResult.list.length"></q-btn>
        </q-item-label>
      </q-item-section>
    </q-item>

    <q-item v-if="authorResult.list.length > 0">
      <q-item-section side top>
        <q-btn flat :color="!form.author?'primary':'black'"  label="作者" @click="handleRemoveQuery('author')"></q-btn>
      </q-item-section>
      <q-item-section>
        <q-item-label>
          <template v-for="(author, index) in authorResult.list" :key="index">
            <q-btn flat
                   :label="author.name"
                   :color="form.author == author.id ?'primary':'black'"
                   @click="handleQuery({author: author.id})" />
          </template>
          <q-btn flat label="更多" to="/authors" v-if="authorResult.pagination.rowsNumber > authorResult.list.length"></q-btn>
        </q-item-label>
      </q-item-section>
    </q-item>

    <q-item v-if="tagResult.list.length > 0">
      <q-item-section side top>
        <q-btn flat :color="!form.tag?'primary':'black'"  label="标签" @click="handleRemoveQuery('tag')"></q-btn>
      </q-item-section>
      <q-item-section>
        <q-item-label>
          <template v-for="(tag, index) in tagResult.list" :key="index">
            <q-btn flat
                   :label="tag.name"
                   :color="form.tag == tag.id ?'primary':'black'"
                   @click="handleQuery({tag: tag.id})" />
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
          <q-input dense outlined label="输入诗词标题" @keyup.enter="handlePoems" v-model="result.form.title">
            <template v-slot:append>
              <q-icon name="search" style="cursor: pointer;" @click="handlePoems" />
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
        <q-item-section class="justify-start">
          <div class="row q-gutter-sm full-width">
            <q-card flat bordered :key="index" v-for="(row, index) in result.list" class="col" style="min-width: 32%;">
              <template v-if="row.title != ''">
                <q-card-section :title="handlePoemTitle(row)" class="row justify-between">
                  <div class="text-h6 ellipsis self-baseline" style="cursor: pointer; max-width: 80%;"
                       @click="handlePoem(row)">
                    <template v-if="row.chapter != ''">{{ row.chapter }} · </template>{{ row.title }}
                  </div>
                  <div class="text-subtitle2 ellipsis text-primary self-baseline" style="cursor: pointer"
                       @click="handleAuthor(row.author)"
                       v-if="row.author">
                    {{ row.author.name }}
                  </div>
                </q-card-section>

                <q-separator inset />
              </template>

              <q-card-section>
                <div class="ellipsis-3-lines" style="white-space: pre-line;">
                  {{ row.content }}
                </div>
              </q-card-section>
            </q-card>
          </div>
        </q-item-section>
      </q-item>

      <q-item v-if="paginationTotal > 1">
        <q-item-section class="items-center">
          <q-pagination direction-links boundary-numbers
                        size="sm"
                        :max="paginationTotal"
                        :max-pages="6"
                        :model-value="result.pagination.page"
                        @update:model-value="handlePagination" />
        </q-item-section>
      </q-item>
    </template>
    <q-inner-loading :showing="result.loading"></q-inner-loading>
  </q-list>
</template>

<script setup>
 import { computed, ref, watch, onMounted } from 'vue';
 import { getCurrentInstance } from 'vue'

 const { proxy } = getCurrentInstance()

 const collectionResult = ref({
     list: [],
     pagination: {
         page: 1,
         rowsPerPage: 0,
         rowsNumber: 0,
     }
 })

 const dynastyResult = ref({
     list: [],
     pagination: {
         page: 1,
         rowsPerPage: 0,
         rowsNumber: 0,
     }
 })

 const authorResult = ref({
     list: [],
     pagination: {
         page: 1,
         rowsPerPage: 0,
         rowsNumber: 0,
     }
 })

 const tagResult = ref({
     list: [],
     pagination: {
         page: 1,
         rowsPerPage: 0,
         rowsNumber: 0,
     }
 })

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
     return Math.ceil(result.value.pagination.rowsNumber / result.value.pagination.rowsPerPage)
 })

 const deepClone = (data) => {
     return JSON.parse(JSON.stringify(data));
 }

 const handlePagination = (page) => {
     result.value.pagination.page = page
     handlePoems()
 }

 const handleQuery = (query) => {
     proxy.$router.push({path: "/poems", query: {...form.value, ...query}})
 }

 const handleRemoveQuery = (key) => {
     let query = {...form.value}
     delete query[key]
     proxy.$router.push({path: "/poems", query: query})
 }

 const handleAuthor = (row) => {
     proxy.$router.push({path: `/authors/${row.id}`})
 }

 const handlePoem = (row) => {
     proxy.$router.push({path: `/poems/${row.id}`})
 }

 const handlePoemTitle = (row) => {
     let title = row.title
     if (row.author) {
         title = title + " 作者:" + row.author.name
     }
     return title
 }

 const handleTags = () => {
     const params = {}
     proxy.$api.get("/api/tags", {
         params: params
     }).then(resp => {
         tagResult.value.list = resp.data.data.list
         tagResult.value.pagination.rowsNumber = resp.data.data.total
         tagResult.value.pagination.rowsPerPage = resp.data.data.limit
     })
 }

 const handleAuthors = () => {
     const params = {}
     if (form.value.dynasty) {
         params.dynasty = form.value.dynasty
     }
     proxy.$api.get("/api/authors", {
         params: params
     }).then(resp => {
         authorResult.value.list = resp.data.data.list
         authorResult.value.pagination.rowsNumber = resp.data.data.total
         authorResult.value.pagination.rowsPerPage = resp.data.data.limit
     })
 }

 const handleDynasties = () => {
     proxy.$api.get("/api/dynasties").then(resp => {
         dynastyResult.value.list = resp.data.data.list
         dynastyResult.value.pagination.rowsNumber = resp.data.data.total
         dynastyResult.value.pagination.rowsPerPage = resp.data.data.limit
     })
 }

 const handleCollections = () => {
     const params = {}
     if (form.value.dynasty) {
         params.dynasty = form.value.dynasty
     }
     proxy.$api.get("/api/collections", {
         params: params
     }).then(resp => {
         collectionResult.value.list = resp.data.data.list
         collectionResult.value.pagination.rowsNumber = resp.data.data.total
         collectionResult.value.pagination.rowsPerPage = resp.data.data.limit
     })
 }

 const handlePoems = () => {
     const params = {...form.value}
     if (result.value.form.title) {
         params.title = result.value.form.title
     }
     if (result.value.pagination.page > 1) {
         params.page = result.value.pagination.page
     }
     result.value.loading = true
     proxy.$api.get("/api/poems", {
         params: params
     }).then(resp => {
         result.value.list = resp.data.data.list
         result.value.pagination.rowsNumber = resp.data.data.total
         result.value.pagination.rowsPerPage = resp.data.data.limit
     }).finally(_ => {
         result.value.loading = false
     })
 }

 watch(() => form.value,
       (newQuery, oldQuery) => {
           result.value.pagination.page = 1
           if (newQuery.dynasty != oldQuery.dynasty) {
               handleAuthors()
               handleCollections()
           }
           handlePoems()
       }
 )

 onMounted(() => {
     handleTags()
     handleAuthors()
     handlePoems()
     handleDynasties()
     handleCollections()
 })
</script>
