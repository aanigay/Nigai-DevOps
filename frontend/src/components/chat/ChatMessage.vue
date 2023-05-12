<template>
    <div>
        <div v-if="name === 'admin'" class="mes">
            <p class="text-xs-center">{{ text }}</p>
            <hr>
        </div>
        <div v-else class="wrap">
            <div class="mes" :class="{owner}">
                <p>
                    {{time}}
                </p>
                <small>
                    <strong>{{ name }}</strong>
                </small>
                <p>{{ text }}</p>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    props: {
        name: String,
        text: String,
        owner: {
            type: Boolean,
            default: false
        },
        sent_at: Number
    },
    computed: {
        time() {
            let date = new Date(this.sent_at)
            let ho = new Intl.DateTimeFormat('ru', { hour: '2-digit' }).format(date);
            let mi = new Intl.DateTimeFormat('ru', { minute: '2-digit' }).format(date);

            return `${ho}:${mi}`
        }
    }
};
</script>

<style scoped>
system {
    margin-bottom: 1rem;

    p {
        margin-bottom: 1rem;
    }
}

.wrap {
    display: flex;
    flex-direction: column;
}

.mes {
    padding: 1rem;
    width: 60%;
    margin: 0 1rem;
    box-shadow: 0 1px 0 0 rgba(50, 50, 50, 0.3);
    border-radius: 4px;
    background: #80CBC4;
    position: relative;
    margin-bottom: 1rem;

    p {
        margin-bottom: 0;
    }
}

.owner {
    background: #eeeeee;
    color: #000000;
    align-self: flex-end;
}
</style>

