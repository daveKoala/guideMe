<template>
    <div class="card">
        <Timeline :value="events" align="left" class="customized-timeline">
            <template #marker="slotProps">
                <span class="timeline-marker" :style="{ backgroundColor: slotProps.item.color }">
                    <i :class="slotProps.item.icon"></i>
                </span>
            </template>
            <template #content="slotProps">
                <Card>
                    <template #title>
                        {{ slotProps.item.status }}
                    </template>
                    <template #subtitle>
                        {{ slotProps.item.date }}
                    </template>
                    <template #content>
                        <img v-if="slotProps.item.image" :src="`https://primefaces.org/cdn/primevue/images/product/${slotProps.item.image}`" :alt="slotProps.item.name" width="200" class="timeline-image" />
                        <p>
                            Lorem ipsum dolor sit amet, consectetur adipisicing elit. Inventore sed consequuntur error repudiandae numquam deserunt quisquam repellat libero asperiores earum nam nobis, culpa ratione quam perferendis esse, cupiditate
                            neque quas!
                        </p>
                        <Button label="Read more" variant="text"></Button>
                    </template>
                </Card>
            </template>
        </Timeline>
    </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import Timeline from "primevue/timeline";
import Card from "primevue/card";
import Button from "primevue/button";

const events = ref([
    { status: 'Ordered', date: '15/10/2020 10:30', icon: 'pi pi-shopping-cart', color: '#9C27B0', image: 'game-controller.jpg' },
    { status: 'Processing', date: '15/10/2020 14:00', icon: 'pi pi-cog', color: '#673AB7' },
    { status: 'Shipped', date: '15/10/2020 16:15', icon: 'pi pi-shopping-cart', color: '#FF9800' },
    { status: 'Delivered', date: '16/10/2020 10:00', icon: 'pi pi-check', color: '#607D8B' }
]);

</script>

<style lang="scss" scoped>
:deep(.p-timeline-event-opposite) {
    flex: 0;
    padding: 0;
}

@media screen and (max-width: 960px) {
    ::v-deep(.customized-timeline) {
        .p-timeline-event:nth-child(even) {
            flex-direction: row;

            .p-timeline-event-content {
                text-align: left;
            }
        }
    }
}

.p-card {
    margin-bottom: 1em;
}

.timeline-marker {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 2rem;
    height: 2rem;
    color: #fff;
    border-radius: 50%;
    z-index: 1;
    box-shadow: 0 1px 2px rgb(0 0 0 / 0.1);
}

.timeline-image {
    margin-bottom: 0.5rem;
    box-shadow: 0 1px 2px rgb(0 0 0 / 0.1);
}
</style>