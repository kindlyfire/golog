<template>
    <ul>
        <router-link tag="li" to="/" exact>
            <a>Dashboard</a>
        </router-link>
        <router-link tag="li" to="/posts" exact>
            <a>Posts</a>
        </router-link>
        <router-link tag="li" to="/media" exact>
            <a>Media</a>
        </router-link>
        <router-link tag="li" to="/pages" exact>
            <a>Pages</a>
        </router-link>
        <router-link tag="li" to="/comments" exact>
            <a>Comments</a>
        </router-link>

        <li class="spacer"></li>

        <router-link tag="li" to="/appearance" exact>
            <a>Appearance</a>
        </router-link>
        <router-link tag="li" to="/plugins" exact>
            <a>Plugins</a>
        </router-link>
        <router-link tag="li" to="/users" exact>
            <a>Users</a>
        </router-link>
        <router-link tag="li" to="/settings" exact>
            <a>Settings</a>
        </router-link>
    </ul>
</template>

<style scoped>
ul {
    list-style-type: none;
    padding: 0;
    padding-top: 75px;
    padding-bottom: 50px;
}
li {
    margin: 0;
    font-size: 1.2em;
    transition: all 0.3s;
    cursor: pointer;
}
li:hover {
    background-color: rgba(173,136,183,0.08);
}
li.spacer {
    height: 30px;
}
a {
    color: inherit;
    display: block;
    outline: none;
    padding: 15px 40px;
}
</style>
