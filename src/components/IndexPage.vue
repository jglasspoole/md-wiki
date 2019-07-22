<template>
  <div class="container">
    <div class="row">
      <div id="index-title" class="col-md-12">
        <h1>Articles Index</h1>
      </div>
    </div>
    <hr />
    <div class="row">
      <div id="article-links" class="col-md-12">
        <div v-for="article in articleTitles" v-bind:key="article">
          <router-link :to="`/${article}`">
            <h4>{{ article }}</h4>
          </router-link>
        </div>
        <div v-if="articleTitles.length == 0 && pageInit" class="alert alert-warning" role="alert">
          No articles found!
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default{
  data () {
    return {
      articleTitles: [],
      pageInit: false // to avoid flashing warning before load complete
    }
  },
  methods: {

  },
  created () {
    this.$http.get('http://localhost:8080/articles').then(function (data) {
      this.articleTitles = data.body.titles
      this.pageInit = true
    })
  }
}
</script>
