import auth from './../constants/auth'

export default [
    {
        name: "Dashboard",
        route: "/",
        iconClass: "",
        allowedRoles: auth.USER
    },
    {
        name: "Report",
        route: "/report",
        iconClass: "",
        allowedRoles: auth.ADMIN
    },
    {
        name: "Settings",
        iconClass: "",
        allowedRoles: auth.ADMIN,
        children: [
            {
                name: "Users",
                route: "/settings/users",
                iconClass: "",
                allowedRoles: auth.ADMIN
            },
            {
                name: "Zones",
                route: "/settings/zones",
                iconClass: "",
                allowedRoles: auth.ADMIN
            },
            {
                name: "Products",
                route: "/settings/products",
                iconClass: "",
                allowedRoles: auth.ADMIN
            }
        ]
    }
];
