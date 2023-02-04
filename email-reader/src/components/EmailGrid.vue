<script setup>
    defineProps(['correos']);
    const emit = defineEmits(['handleClick']);

    const formatDate = (date = '') => {
        const timestamp = Date.parse(date);
        const formattedDate = new Date(timestamp).toUTCString();
        return formattedDate;
    }
</script>

<template>
    <table class="table-auto w-full divide-y-2 divide-gray">
        <thead class="pt-5 h-11">
            <tr class="">
                <th class="sticky top-0 bg-secundario text-left px-5 font-semibold text-gray">Subject</th>
                <th class="sticky top-0 bg-secundario text-left px-5 font-semibold text-gray">From</th>
                <th class="sticky top-0 bg-secundario text-left px-5 font-semibold text-gray">Date</th>
            </tr>
        </thead>
        <tbody class="">
            <tr
                class="cursor-pointer hover:bg-blueButtons shadow-2xl divide-x divide-gray"
                :class="idx % 2 === 0 ? 'bg-primario' : 'bg-secundario'"
                v-for="correo, idx in correos"
                :key="correo.id"
                @click="emit('handleClick', correo.id)"
            >
                <td class="px-5 py-1 text-white">{{correo.subject}}</td>
                <td class="px-5 py-1 text-white">{{correo.from}}</td>
                <td class="px-5 py-1 text-white">{{formatDate(correo.date)}}</td>
            </tr>
        </tbody>
    </table>
</template>