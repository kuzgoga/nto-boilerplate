<script setup lang="ts">
import { ref, onMounted } from "vue";
import { Events } from "@wailsio/runtime";
import * as PostService from "../../bindings/app/internal/services/postservice";

const name = ref("");
const result = ref("Please enter your name below 👇");
const time = ref("Listening for Time event...");

const doGreet = () => {
    let localName = name.value;
    if (!localName) {
        localName = "anonymous";
    }
};

onMounted(async () => {
    console.log(await PostService.GetById(5));
});
</script>

<template>
    <h1>Kuzbass</h1>

    <div class="result"></div>
    <div class="card">
        <div class="input-box">
            <input
                class="input"
                v-model="name"
                type="text"
                autocomplete="off"
            />
            <button class="btn" @click="doGreet">Greet</button>
        </div>
    </div>

    <div class="footer">
        <div><p>Click on the Wails logo to learn more</p></div>
        <div><p></p></div>
    </div>
</template>
