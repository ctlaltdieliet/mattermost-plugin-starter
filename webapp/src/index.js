import React from 'react';

// Courtesy of https://feathericons.com/
const Icon = () => <i className='icon fa fa-plug'/>;

class HelloWorldPlugin {
    initialize(registry, store) {
        
      
    }
}

window.registerPlugin('be.controlaltdieliet.demoortom.first_mattermost_plugin', new HelloWorldPlugin());
