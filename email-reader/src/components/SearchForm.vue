<script setup>
    import { ref } from 'vue';

    const emit = defineEmits(['handleSubmit']);

    const field = ref('');
    const query = ref('');

    const fields = [
        {name: 'All', value: ''},
        {name: 'Subject', value: 'subject'},
        {name: 'From', value: 'from'},
        {name: 'To', value: 'to'},
        {name: 'Date', value: 'date'},
        {name: 'Body', value: 'body'},
    ]

    const handleClick = (field, query) => {
        if (field) {
            return emit('handleSubmit', { field, query });
        }

        emit('handleSubmit', { query });
    }

    const handleKey = (e) => {
        if (e.keyCode === 13) {
            handleClick(field.value, query.value);
        }
    };

    const handleSelect = (e) => {
        if (!query.value) {
            return;
        }

        emit('handleSubmit', {query, field})
    }

</script>

<template>

    <div class="w-full flex shadow-2xl">
        <div class="bg-blueButtons px-5 rounded-l-lg shadow-2xl">
            <select 
                v-model="field" 
                class="appearance-none my-auto w-full h-full focus:outline-none bg-blueButtons text-white cursor-pointer"
                @change="handleSelect"
            >
                <option v-for="field in fields"
                    :value="field.value"
                >
                    {{field.name}}
                </option>
            </select>
        </div>

        <input
            class="grow h-11 px-3 py-1 focus:outline-none focus:border-2 focus:border-mostaza"
            v-model="query" 
            placeholder="Que deseas buscar" 
            @keypress="handleKey"
        />

        <button 
            class="bg-mostaza px-5 py-1 border border-mostaza h-11 text-white rounded-r-lg hover:bg-primario hover:border-2 transition-all ease-in-out duration-300"
            @click="() => handleClick(field, query)"
        >
            Search
        </button>
    </div>

</template>