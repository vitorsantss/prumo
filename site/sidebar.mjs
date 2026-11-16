export const menu = [
  {
    label: 'Comece',
    translations: { en: 'Start' },
    items: ['docs/comece-aqui', 'docs/primeira-contribuicao'],
  },
  {
    label: 'Entenda',
    translations: { en: 'Understand' },
    items: [
      'docs/a-ideia',
      'docs/processo',
      { label: 'Decisões (ADRs)', translations: { en: 'Decisions (ADRs)' }, collapsed: true, items: [{ autogenerate: { directory: 'docs/adr' } }] },
    ],
  },
  {
    label: 'Faça',
    translations: { en: 'Do' },
    items: ['docs/formas-de-contribuir', 'docs/pruminhos', 'docs/traduzir', 'docs/como-mudar-o-site'],
  },
  {
    label: 'Consulte',
    translations: { en: 'Look up' },
    items: ['docs/estado-atual', 'docs/publicacao', 'docs/privacidade', 'docs/referencias', { label: 'Design system', slug: 'docs/design' }, 'docs'],
  },
  {
    label: 'Design system',
    items: [
      { label: 'Visão geral', translations: { en: 'Overview' }, slug: 'docs/design' },
      {
        label: 'Fundamentos',
        translations: { en: 'Foundations' },
        items: ['docs/design/fundamentos/principios', 'docs/design/fundamentos/cor', 'docs/design/fundamentos/tipografia', 'docs/design/fundamentos/espaco-e-raio', 'docs/design/fundamentos/movimento', 'docs/design/fundamentos/acessibilidade'],
      },
      {
        label: 'Componentes',
        translations: { en: 'Components' },
        items: [
          { label: 'Site', items: ['docs/design/componentes/seletor-de-idioma'] },
          { label: 'Os dois', translations: { en: 'Both' }, items: ['docs/design/componentes/botao', 'docs/design/componentes/ponto-de-status'] },
        ],
      },
      { label: 'Regras da camada Site', translations: { en: 'Site layer rules' }, slug: 'docs/design/site' },
      'docs/design/como-propor-um-componente',
      { label: 'Voltar pra documentação', translations: { en: 'Back to the docs' }, slug: 'docs/comece-aqui' },
    ],
  },
];
