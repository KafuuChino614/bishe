<template>
  <div class="flex gap-4 items-center text-sm text-gray-500 justify-center mb-4">
    <span>Powered by</span>
    <span>
      <a
        class="text-gray-800 font-bold"
      >HAUT</a>
    </span>
    <slot />
    <span>Copyright</span>
    <span>
      <a
        class="text-gray-800 font-bold"
      >lyk</a>
    </span>
  </div>
</template>

<script  setup>

defineOptions({
  name: 'BottomInfo'
})

console.log(
  `%c powered by %c flipped-aurorae %c`,
  'background:#0081ff; padding: 1px; border-radius: 3px 0 0 3px; color: #fff',
  'background:#354855; padding: 1px 5px; border-radius: 0 3px 3px 0; color: #fff; font-weight: bold;',
  'background:transparent'
)
</script>

