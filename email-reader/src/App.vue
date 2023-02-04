<script setup>
  import { computed, onMounted, ref } from 'vue';
  import EmailGrid from './components/EmailGrid.vue';
  import SearchForm from './components/SearchForm.vue'

  //PARAMETROS DE LA APP
  const skip = 20;

  //VARIABLES REACTIVAS DE LA PÁGINA WEB
  const endpoint = ref('')
  const correos = ref(new Map());
  const correo = ref({});
  const showCorreo = ref(false);
  const errorMsg = ref('');
  const searchResultsTotal = ref(0);
  const query = ref({});

  //ON MOUNTED HOOK
  onMounted(() => {
    endpoint.value = window.location.protocol + '//' + window.location.host;
    errorMsg.value = 'Type your search!!!';
  });

  //MÉTODOS COMPUTADOS
  const correosArr = computed(() => {
    return Array.from(correos.value.values());
  })

  // MANEJO EN EL CLICK DE CORREO SELECCIONADO EN LA LISTA
  const handleClick = (id) => {

    if (correo.value.id === id && showCorreo.value) {
      return showCorreo.value = false;
    }

    correo.value = correos.value.get(id);
    showCorreo.value = true;
  }

  //MANEJO DE BOTONES DE PAGINADO
  const handlePrev = () => {
    if (query.value.range[0] === 0) {
      return;
    } 

    const newRange = [
      query.value.range[0] - skip,
      query.value.range[1] - skip
    ] 

    query.value.range = newRange;
    handleSubmit(query.value);
  }

  const handleNext = () => {

    if (query.value.range[1] >= searchResultsTotal.value) {
      return;
    }

    const newRange = [
      query.value.range[0] + skip,
      query.value.range[1] + skip
    ] 

    query.value.range = newRange;
    handleSubmit(query.value);
  }


  //MANEJO DEL BOTÓN DE BUSQUEDA
  const handleSubmit = async (jsonQuery = {}) => {

    //Validación para ver si se página en primera página
    if (query.value !== jsonQuery) {
      query.value = jsonQuery;
      query.value.range = [0, skip];
    }

    correos.value.clear();
    showCorreo.value = false;

    const url = endpoint.value + "/api/busqueda"

    const headers = new Headers({
      "Content-Type": "application/json",
      "Authorization": "YWRtaW46Q29tcGxleHBhc3MjMTIz"
    });

    const body = JSON.stringify(query.value);

    try {
      const res = await fetch(url, {
        method: 'POST',
        headers,
        body
      });

      const { results = [], msg = '', total = 0 } = await res.json();

      if (msg) {
        return errorMsg.value = 'Empty query, please try again!';
      } else if (total === 0) {
        return errorMsg.value = 'No results.';
      } else {
        errorMsg.value = '';
      }

      results.forEach(result => correos.value.set(result.id, result));
      searchResultsTotal.value = total;


    } catch (err) {
      console.log(err);
      return errorMsg.value = 'Connection lost, please try again!';
    }

  }
</script>

<template>

  <!-- CONTENEDOR -->
  <div class="p-5">

    <!-- COMPONENTE DE BARRA DE BUSQUEDA -->
    <div class="p-5">
      <SearchForm
        @handle-submit="handleSubmit"
      />
    </div>


    <!-- CONTENEDOR DE VISUALIZADOR DE RESULTADOS Y LECTOR -->
    <div 
      v-if="!errorMsg && correosArr.length > 0"
      class="grid grid-cols-2 p-5 m-5 gap-5 bg-primario rounded-lg shadow-2xl"
    >

      <!-- TOTAL -->
      <div class="px-5 col-span-full flex justify-start">
        <p class="text-white">
          {{searchResultsTotal}} search results
        </p>
      </div>
      <!-- TOTAL FIN -->

      <!-- LISTADO DE CORREOS -->
      <div 
        class="p-5 pt-0 overflow-auto h-[33rem]"
        :class="!showCorreo ? 'col-span-full' : null"
      >
        <EmailGrid
          :correos="correosArr"
          @handle-click="handleClick"
        />
      </div>
      <!-- LISTADO DE CORREOS FIN -->

      <!-- VISUALIZADOR DE CORREO -->
      <div v-show="showCorreo" class="h-[33rem] overflow-auto bg-gray shadow-2xl rounded-lg p-5">
        <div class="overflow-auto">
          <h2 class="whitespace-pre">{{correo.subject}}</h2>
          <p>From: {{correo.from}}</p>
          <br/>
          <p class="whitespace-pre">
            {{correo.body}}
          </p>
        </div>
      </div>
      <!-- VISUALIZADOR FIN -->

      <!-- COMPONENTE DE PAGINACIÓN -->
      <div class="col-span-full px-5">
        <div class="flex justify-between">
          <button 
            @click="handlePrev"
            class="text-white bg-blueButtons px-5 py-3 rounded-lg"
          >
            Prev Page
          </button>
          <p class="text-gray my-auto">
            {{query.range[0] + 1}} to {{query.range[1] > searchResultsTotal ? searchResultsTotal : query.range[1]}} of {{searchResultsTotal}}
          </p>
          <button 
            @click="handleNext"
            class="text-white bg-blueButtons px-5 py-3 rounded-lg"
          >
            Next Page
          </button>
        </div>
      </div>

    </div>
    <!-- VISUALIZADOR Y LECTOR FIN -->

    <!-- DISPLAY DE ERRORES -->
    <div 
      v-else
      class="px-5 mt-5 w-full flex justify-center"
    >
      <p class="text-red-500 font-semibold text-lg">
        {{ errorMsg }}
      </p>
    </div>
    <!-- DISPLAY DE ERRORES FIN -->

  </div>
  <!-- CONTENEDOR FIN -->

</template>