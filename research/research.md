## 1️⃣ Introdução profissional

> Ao longo da minha carreira de mais de 8 anos como engenheiro de software, desenvolvi especialização sólida em **Golang**, que se tornou a principal linguagem no meu dia a dia, especialmente em projetos que exigem alta performance, escalabilidade e resiliência.
>
> Tenho experiência profunda com AWS, construindo soluções escaláveis, observáveis e cloud-native, aplicando boas práticas de microsserviços, containers e infraestrutura como código. Também utilizo inteligência artificial como ferramenta no meu fluxo de trabalho, sempre buscando formas de aumentar produtividade e qualidade.
>
> Além da minha proficiência técnica, sou muito cuidadoso em escrever código limpo, testável, com foco na legibilidade e manutenção, valorizando princípios como os **12 fatores para aplicações modernas e os conceitos fundamentais de design e performance**.
>
> Embora o meu foco recente tenha sido majoritariamente em Golang, tenho experiência prévia com linguagens estruturadas e orientadas a objeto como Java e C.
> Como profissional, entendo que a linguagem é uma ferramenta para resolver problemas e, por isso, mantenho a flexibilidade e disciplina para reaprender rapidamente tecnologias e stacks que não utilizo no dia a dia.
>
> Minha experiência em ambientes ágeis, somada à minha capacidade de adaptação, me permite retomar o uso de Java ou C sempre que o projeto exigir — com a mesma atenção à qualidade, performance e padrões de engenharia que aplico hoje em Go.
>
> Estou confiante que essa combinação de profundidade em Golang, conhecimento abrangente de boas práticas de engenharia e habilidade para reaprender outras linguagens com rapidez vai permitir que eu atue de forma eficaz e colaborativa com o time da NTT DATA e Itaú Banco.

---

## 2️⃣ Comparação Go vs C vs Java no contexto Layer 0 bancário

| Aspecto                   | **Go**                                                  | **C**                              | **Java**                              |
| ------------------------- | ------------------------------------------------------- | ---------------------------------- | ------------------------------------- |
| **Compilação**            | Binário nativo, estático, ideal para ambientes críticos | Binário nativo, toolchain complexo | Bytecode + JVM (maior footprint)      |
| **Gerenciamento memória** | Garbage Collector moderno, tuning simples               | Manual, mais risco de erro humano  | GC da JVM, tuning avançado necessário |
| **Concorrência**          | Goroutines leves e eficientes                           | pthreads, manual e verboso         | Threads pesadas, precisa frameworks   |
| **Deploy**                | Binário único e portátil                                | Binário dependente do ambiente     | JVM + dependências                    |
| **Performance**           | Alta, próxima de C                                      | Máxima, mas manutenção difícil     | Boa, overhead JVM                     |
| **Complexidade**          | Minimalista, legível                                    | Complexo e propenso a bugs         | Verboso, frameworks pesados           |
| **Fit Layer 0 Bancário**  | Ideal para APIs críticas e confiáveis                   | Adequado para baixo nível          | Bom para sistemas legados             |

✅ Vantagens práticas de Go no Layer 0 bancário:

- Performance + concorrência eficiente
- Deploy simples e previsível
- Menos complexidade
- Fácil integração com containers e cloud-native infra

---

## 3️⃣ 12 princípios de software (12-Factor App) aplicados ao Layer 0 bancário

**Resumo:** _Conjunto de práticas modernas que garantem portabilidade, escalabilidade, resiliência e facilidade de manutenção para aplicações cloud-ready e de missão crítica._

1️⃣ Base de código única: controle rigoroso, rastreável, compliance-ready
2️⃣ Dependências declaradas: builds previsíveis e auditáveis
3️⃣ Configuração externa: segredos e configs seguros via env vars e vaults
4️⃣ Serviços de apoio como recursos: fácil substituição sem acoplamento rígido
5️⃣ Build, release, run distintos: ciclos consistentes, rollback-friendly
6️⃣ Stateless: essencial para escalabilidade horizontal e alta disponibilidade
7️⃣ Binding de porta: APIs self-contained, ideais para ambientes controlados
8️⃣ Escala horizontal: processos replicáveis suportando volumes críticos
9️⃣ Start/shutdown rápido: fundamental para resiliência e failover
🔟 Ambientes homogêneos: redução de risco de bugs/erros em produção
⓫ Logs como stream: integração nativa com plataformas de observabilidade
⓬ Admin processes: scripts isolados para manutenção confiável e auditável
⓬ Admin processes isolados

---

## 4️⃣ Extras com sub-tópicos no contexto Layer 0

### 🔹 Databases: estratégia típica no Layer 0 bancário

- **Prioridade:** confiabilidade, consistência e compliance regulatório.
- **Estratégias comuns:**

  - Escalabilidade vertical com clustering ativo-passivo.
  - Read replicas para relatórios e analytics.
  - Horizontal scaling muito restrito.
  - Tecnologias: Oracle RAC, IBM DB2, SQL Server Always On, PostgreSQL clustering.

- **Princípio central:** evitar arquiteturas distribuídas excessivas para preservar integridade e auditabilidade.

### 🔹 ACID: resumo no contexto bancário

- **Atomicity:** tudo-ou-nada, sem estados intermediários.
- **Consistency:** mantém integridade e regras de negócio.
- **Isolation:** concorrência segura.
- **Durability:** persistência garantida mesmo em falhas.

✅ Crítico no Layer 0: confiança, auditoria e integridade transacional.

### 🔹 Kafka: pequeno resumo (Layer 0)

- Integração assíncrona confiável e escalável.
- Resiliência e ordering garantido.
- Ideal para event-driven architectures críticas.

### 🔹 Kubernetes, EKS, AKS, GKE (Layer 0)

- Orquestração moderna com alta disponibilidade.
- EKS/AKS/GKE: gerenciados, reduzindo complexidade operacional.

### 🔹 Terraform / CloudFormation (Layer 0)

- **Terraform:** multi-cloud, declarativo, controle versionado.
- **CloudFormation:** AWS-native, integração profunda com AWS.

---

💡 Todo o conteúdo está contextualizado para Layer 0 bancário: performance, resiliência, auditabilidade e alta disponibilidade.
