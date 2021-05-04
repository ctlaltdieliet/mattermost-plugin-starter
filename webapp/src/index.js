import React from 'react';

// Courtesy of https://feathericons.com/
const Icon = () => <i className='icon fa fa-plug'/>;

class HelloWorldPlugin {
    initialize(registry, store) {
        
        registry.registerMainMenuAction(
            'Plugin Menu Item',
            () => console.log("Hi"),
         Icon
        );
        registry.registerChannelHeaderButtonAction(
            <Icon />,
            () => {
                alert("Hello there");
            },
            "Hello you",
        );
    }
}

window.registerPlugin('be.controlaltdieliet.demoortom.first_mattermost_plugin', new HelloWorldPlugin());
