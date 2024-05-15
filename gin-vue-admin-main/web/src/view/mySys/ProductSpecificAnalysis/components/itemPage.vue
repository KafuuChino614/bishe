 <template>
  <div class="item">
      <!-- 设置插槽 -->
      <slot></slot>
  </div>
</template>

<script>
export default {

}
</script>

<style lang="less">
.item{
    height: 700px;
    border: 1px solid blue;
    margin: 20px;
    background-color: rgba(12,130,255,.85);
}
</style>