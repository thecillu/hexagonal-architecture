package com.cillu;

import com.cillu.adapters.MemoryRepository;
import com.cillu.core.services.TeamService;
import com.cillu.ports.RepositoryPort;
import com.cillu.ports.TeamServicePort;
import org.springframework.context.annotation.Bean;

@org.springframework.context.annotation.Configuration
public class Configuration {

    @Bean
    public RepositoryPort repositoryPort() {
        return new MemoryRepository();
    }


@Bean
public TeamServicePort teamServicePort() {
    return new TeamService(repositoryPort());
}

}
